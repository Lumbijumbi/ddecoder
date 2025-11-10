Encodes and Decodes Payloads made by Datadome Anti Bot.

This software is provided for educational and research purposes only.

## Components

- **v1/**: Core encoder/decoder library (Go package)
- **cmd/gui/**: Standalone GUI application with macOS support for encoding and decoding

## Using the Library

The encoder/decoder can be used as a Go library:

```go
import decoder "github.com/sting420/ddecoder/v1"

// Decode a payload
decoded, err := decoder.Decode(data, decoder.Interstitial)

// Encode JSON to payload
encoded, err := decoder.Encode(jsonData, cid, hash, decoder.Interstitial)
```

## Standalone GUI

A cross-platform GUI application is available in `cmd/gui/` with support for both encoding and decoding. See [cmd/gui/README.md](cmd/gui/README.md) for build instructions and usage.

### Quick Start (macOS)

```bash
cd cmd/gui
make build
./ddecoder-gui
```