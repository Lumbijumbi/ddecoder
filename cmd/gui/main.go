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
	w := a.NewWindow("DataDome Payload Decoder")
	w.Resize(fyne.NewSize(800, 600))

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
	w.SetContent(container.NewBorder(
		container.NewVBox(
			widget.NewLabel("DataDome Payload Decoder"),
			widget.NewSeparator(),
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
				widget.NewLabel("Output (Decoded Data):"),
				nil, nil, nil,
				output,
			),
		),
	))

	w.ShowAndRun()
}
