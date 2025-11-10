package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	decoder "github.com/sting420/ddecoder/v1"
)

func main() {
	a := app.New()
	w := a.NewWindow("DataDome Payload Encoder/Decoder")
	w.Resize(fyne.NewSize(900, 650))

	// Create decoder tab
	decoderTab := createDecoderTab()

	// Create encoder tab
	encoderTab := createEncoderTab()

	// Create tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Decoder", decoderTab),
		container.NewTabItem("Encoder", encoderTab),
	)

	w.SetContent(tabs)
	w.ShowAndRun()
}

func createDecoderTab() *fyne.Container {
	// Input field for encoded data
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Paste encoded data here...")
	input.Wrapping = fyne.TextWrapWord

	// Output field for decoded data
	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("Decoded result will appear here...")
	output.Wrapping = fyne.TextWrapWord
	output.Disable()

	// Seed type selector
	seedSelect := widget.NewSelect([]string{"Interstitial", "Captcha", "Tags"}, nil)
	seedSelect.SetSelected("Interstitial")

	// Decode button
	decodeButton := widget.NewButton("Decode", func() {
		data := input.Text
		if data == "" {
			output.SetText("Error: Please provide input data")
			return
		}

		var seed decoder.EncodingSeed
		switch seedSelect.Selected {
		case "Interstitial":
			seed = decoder.Interstitial
		case "Captcha":
			seed = decoder.Captcha
		case "Tags":
			seed = decoder.Tags
		default:
			seed = decoder.Interstitial
		}

		result, err := decoder.Decode(data, seed)
		if err != nil {
			output.SetText("Error: " + err.Error())
			return
		}

		output.SetText(result)
	})

	// Clear button
	clearButton := widget.NewButton("Clear", func() {
		input.SetText("")
		output.SetText("")
	})

	// Layout
	return container.NewBorder(
		container.NewVBox(
			container.NewHBox(
				widget.NewLabel("Seed Type:"),
				seedSelect,
			),
		),
		container.NewHBox(decodeButton, clearButton),
		nil,
		nil,
		container.NewVSplit(
			container.NewBorder(
				widget.NewLabel("Input (Encoded Data):"),
				nil, nil, nil,
				input,
			),
			container.NewBorder(
				widget.NewLabel("Output (Decoded JSON):"),
				nil, nil, nil,
				output,
			),
		),
	)
}

func createEncoderTab() *fyne.Container {
	// Input field for JSON data
	jsonInput := widget.NewMultiLineEntry()
	jsonInput.SetPlaceHolder("Paste JSON data here...")
	jsonInput.Wrapping = fyne.TextWrapWord

	// Output field for encoded data
	encodedOutput := widget.NewMultiLineEntry()
	encodedOutput.SetPlaceHolder("Encoded result will appear here...")
	encodedOutput.Wrapping = fyne.TextWrapWord
	encodedOutput.Disable()

	// Seed type selector
	seedSelect := widget.NewSelect([]string{"Interstitial", "Captcha", "Tags"}, nil)
	seedSelect.SetSelected("Interstitial")

	// CID input
	cidInput := widget.NewEntry()
	cidInput.SetPlaceHolder("Enter CID...")

	// Hash input
	hashInput := widget.NewEntry()
	hashInput.SetPlaceHolder("Enter hash...")

	// Encode button
	encodeButton := widget.NewButton("Encode", func() {
		jsonData := jsonInput.Text
		cid := cidInput.Text
		hash := hashInput.Text

		if jsonData == "" || cid == "" || hash == "" {
			encodedOutput.SetText("Error: Please provide JSON data, CID, and hash")
			return
		}

		var seed decoder.EncodingSeed
		switch seedSelect.Selected {
		case "Interstitial":
			seed = decoder.Interstitial
		case "Captcha":
			seed = decoder.Captcha
		case "Tags":
			seed = decoder.Tags
		default:
			seed = decoder.Interstitial
		}

		result, err := decoder.Encode(jsonData, cid, hash, seed)
		if err != nil {
			encodedOutput.SetText("Error: " + err.Error())
			return
		}

		encodedOutput.SetText(result)
	})

	// Clear button
	clearButton := widget.NewButton("Clear", func() {
		jsonInput.SetText("")
		cidInput.SetText("")
		hashInput.SetText("")
		encodedOutput.SetText("")
	})

	// Layout
	return container.NewBorder(
		container.NewVBox(
			container.NewHBox(
				widget.NewLabel("Seed Type:"),
				seedSelect,
			),
			container.NewHBox(
				widget.NewLabel("CID:"),
				cidInput,
			),
			container.NewHBox(
				widget.NewLabel("Hash:"),
				hashInput,
			),
		),
		container.NewHBox(encodeButton, clearButton),
		nil,
		nil,
		container.NewVSplit(
			container.NewBorder(
				widget.NewLabel("Input (JSON Data):"),
				nil, nil, nil,
				jsonInput,
			),
			container.NewBorder(
				widget.NewLabel("Output (Encoded Payload):"),
				nil, nil, nil,
				encodedOutput,
			),
		),
	)
}
