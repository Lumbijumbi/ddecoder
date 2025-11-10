Decodes Payloads made by Datadome Anti Bot.

This software is provided for educational and research purposes only.

## Components

- **v1/**: Core decoder library (Go package)
- **cmd/gui/**: Standalone GUI application with macOS support

## Using the Library

The decoder can be used as a Go library:

```go
import decoder "github.com/sting420/ddecoder/v1"

// Decode a payload
result, err := decoder.Decode(data, decoder.Interstitial)
```

## Standalone GUI

A cross-platform GUI application is available in `cmd/gui/`. See [cmd/gui/README.md](cmd/gui/README.md) for build instructions and usage.

### Quick Start (macOS)

```bash
cd cmd/gui
make build
./ddecoder-gui
```