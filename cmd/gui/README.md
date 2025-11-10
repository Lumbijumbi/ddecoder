# DataDome Encoder/Decoder GUI

A standalone GUI application for encoding and decoding DataDome payloads, built with [Fyne](https://fyne.io) for cross-platform compatibility including macOS.

## Features

- User-friendly graphical interface with tabbed interface
- **Decoder Tab**: Decode DataDome payloads to JSON
- **Encoder Tab**: Encode JSON data to DataDome payloads
- Support for all encoding seed types (Interstitial, Captcha, Tags)
- Cross-platform support (macOS, Windows, Linux)
- Real-time encoding/decoding with error handling
- Native macOS app bundle support
- High-resolution display (Retina) support on macOS
- Follows macOS Human Interface Guidelines

## macOS Compatibility

This GUI application is fully compatible with macOS and includes:

- **Native macOS app bundle**: Can be packaged as a `.app` for easy installation
- **Info.plist metadata**: Proper macOS application metadata
- **High DPI support**: Optimized for Retina displays
- **Minimum macOS version**: 10.13 (High Sierra) and later
- **Fyne framework**: Uses native macOS rendering for optimal performance
- **Code signing ready**: Structure supports Apple Developer signing

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

### Quick Build with Make

```bash
cd cmd/gui
make build
./ddecoder-gui
```

### Manual Build

```bash
cd cmd/gui
go build -o ddecoder-gui
./ddecoder-gui
```

### Creating a macOS App Bundle

To create a proper macOS application bundle:

```bash
# Install fyne CLI tool
go install fyne.io/fyne/v2/cmd/fyne@latest

# Package as macOS app (uses Info.plist and FyneApp.toml)
fyne package -os darwin

# Or use Make
make package-macos
```

This will create `DataDome Decoder.app` that can be:
- Moved to your Applications folder
- Launched from Spotlight
- Added to the Dock
- Optionally code-signed for distribution

## Building on Other Platforms

### Windows

```bash
cd cmd/gui
go build -o ddecoder-gui.exe
```

Or create a Windows app:
```bash
make package-windows
```

### Linux

```bash
cd cmd/gui
go build -o ddecoder-gui
```

Or create a Linux package:
```bash
make package-linux
```

## Usage

### Decoder Tab

1. Launch the application and select the "Decoder" tab
2. Select the appropriate seed type (Interstitial, Captcha, or Tags)
3. Paste your encoded DataDome payload in the input field
4. Click "Decode" to see the decoded JSON result
5. Use "Clear" to reset the fields

### Encoder Tab

1. Launch the application and select the "Encoder" tab
2. Select the appropriate seed type (Interstitial, Captcha, or Tags)
3. Enter the CID (Client ID) value
4. Enter the hash value
5. Paste your JSON data in the input field
6. Click "Encode" to see the encoded payload result
7. Use "Clear" to reset all fields

## Seed Types

- **Interstitial**: For interstitial page payloads
- **Captcha**: For CAPTCHA-related payloads
- **Tags**: For tag-based payloads

## Architecture

The GUI application uses:
- **Fyne v2**: Modern, native-looking GUI framework
- **Go modules**: For dependency management
- **Local encoder/decoder library**: Uses the core ddecoder library via Go replace directive
- **Tabbed interface**: Separate tabs for encoding and decoding operations

## License

This software is provided for educational and research purposes only.
