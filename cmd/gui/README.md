# DataDome Decoder GUI

A standalone GUI application for decoding DataDome payloads, built with [Fyne](https://fyne.io) for cross-platform compatibility including macOS.

## Features

- User-friendly graphical interface
- Support for all encoding seed types (Interstitial, Captcha, Tags)
- Cross-platform support (macOS, Windows, Linux)
- Real-time decoding with error handling

## Prerequisites

- Go 1.24 or later
- For macOS: Xcode Command Line Tools

## Building on macOS

### Install Prerequisites

1. Install Xcode Command Line Tools (if not already installed):
```bash
xcode-select --install
```

2. Install Go from [golang.org](https://golang.org/dl/) or using Homebrew:
```bash
brew install go
```

### Build the Application

```bash
cd cmd/gui
go build -o ddecoder-gui
```

### Run the Application

```bash
./ddecoder-gui
```

### Creating a macOS App Bundle

To create a proper macOS application bundle:

```bash
# Install fyne CLI tool
go install fyne.io/fyne/v2/cmd/fyne@latest

# Package as macOS app
fyne package -os darwin -icon Icon.png
```

This will create `ddecoder-gui.app` that can be moved to your Applications folder.

## Building on Other Platforms

### Windows

```bash
cd cmd/gui
go build -o ddecoder-gui.exe
```

Or create a Windows app:
```bash
fyne package -os windows -icon Icon.png
```

### Linux

```bash
cd cmd/gui
go build -o ddecoder-gui
```

## Usage

1. Launch the application
2. Select the appropriate seed type (Interstitial, Captcha, or Tags)
3. Paste your encoded DataDome payload in the input field
4. Click "Decode" to see the decoded result
5. Use "Clear" to reset the fields

## Seed Types

- **Interstitial**: For interstitial page payloads
- **Captcha**: For CAPTCHA-related payloads
- **Tags**: For tag-based payloads

## License

This software is provided for educational and research purposes only.
