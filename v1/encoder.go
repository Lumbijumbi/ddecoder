package decoder

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// calcShift is the inverse of reverseCalcShift
func calcShift(x int32) byte {
	if x == 0 {
		return 45 // corresponds to (45 - 45) / 50 = 0
	} else if x == 1 {
		return 95
	} else if x >= 2 && x <= 11 {
		return byte(x + 46) // x = val - 46, so val = x + 46
	} else if x >= 12 && x <= 37 {
		return byte(x + 53) // x = val - 53, so val = x + 53
	} else { // x > 37
		return byte(x + 59) // x = val - 59, so val = x + 59
	}
}

// combineChunk is the inverse of splitChunk
func combineChunk(bytes []int32) int32 {
	if len(bytes) < 3 {
		return 0
	}
	return (bytes[0] << 16) | (bytes[1] << 8) | bytes[2]
}

// EncodePayload encodes a JSON string into a DataDome payload
func EncodePayload(jsonData string, cid string, hash string, scriptSeed int32) (string, error) {
	var seed int32 = int32(simpleHash(cid)) ^ int32(simpleHash(hash)) ^ scriptSeed

	// Parse JSON to get key-value pairs
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &data); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	// Build the raw string representation as it would be decoded
	var rawStr strings.Builder
	rawStr.WriteString("{")

	first := true
	for key, value := range data {
		if !first {
			rawStr.WriteString(",")
		}
		first = false

		// Format the value based on type
		var valueStr string
		switch v := value.(type) {
		case string:
			valueStr = fmt.Sprintf(`"%s"`, v)
		case float64:
			if v == float64(int64(v)) {
				valueStr = fmt.Sprintf("%d", int64(v))
			} else {
				valueStr = fmt.Sprintf("%v", v)
			}
		case bool:
			valueStr = fmt.Sprintf("%v", v)
		default:
			valueStr = fmt.Sprintf(`"%v"`, v)
		}

		rawStr.WriteString(fmt.Sprintf(`"%s":%s`, key, valueStr))
	}
	rawStr.WriteString("}")

	// Convert to byte array
	rawBytes := []byte(rawStr.String())

	// First step: XOR with seed (reverse of secondStepArray)
	firstStepArray := make([]int32, len(rawBytes))
	for i := 0; i < len(rawBytes); i++ {
		var shift int32 = int32(16 - ((i % 3) * 8))
		firstStepArray[i] = ((seed >> shift) & 255) ^ int32(rawBytes[i])
		if (i % 3) == 2 {
			seed ^= seed << 13
			seed ^= seed >> 17
			seed ^= seed << 5
		}
	}

	// Combine into chunks (reverse of splitChunk)
	var payload strings.Builder
	for i := 0; i < len(firstStepArray); i += 3 {
		chunk := combineChunk(firstStepArray[i:])

		// Encode chunk as 4 characters using calcShift
		b0 := calcShift((chunk >> 18) & 0x3F)
		b1 := calcShift((chunk >> 12) & 0x3F)
		b2 := calcShift((chunk >> 6) & 0x3F)
		b3 := calcShift(chunk & 0x3F)

		payload.WriteByte(b0)
		payload.WriteByte(b1)
		payload.WriteByte(b2)
		payload.WriteByte(b3)
	}

	return payload.String(), nil
}

// Encode creates a complete encoded URL with payload, cid, and hash
func Encode(jsonData string, cid string, hash string, seed EncodingSeed) (string, error) {
	payload, err := EncodePayload(jsonData, cid, hash, int32(seed))
	if err != nil {
		return "", err
	}

	// Build the appropriate URL format based on seed type
	var result string
	switch seed {
	case Captcha:
		result = fmt.Sprintf("ddCaptchaEncodedPayload=%s&hash=%s&cid=%s",
			url.QueryEscape(payload), hash, cid)
	case Interstitial:
		result = fmt.Sprintf("payload=%s&hash=%s&cid=%s",
			payload, hash, cid)
	case Tags:
		result = fmt.Sprintf("jspl=%s&ddk=%s&cid=%s",
			payload, hash, cid)
	default:
		result = fmt.Sprintf("payload=%s&hash=%s&cid=%s",
			payload, hash, cid)
	}

	return result, nil
}
