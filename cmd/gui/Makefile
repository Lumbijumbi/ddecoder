.PHONY: build clean run package-macos package-windows package-linux install-deps

# Build the GUI application
build:
	go build -o ddecoder-gui .

# Run the application
run:
	go run .

# Clean build artifacts
clean:
	rm -f ddecoder-gui ddecoder-gui.exe
	rm -rf ddecoder-gui.app

# Install dependencies
install-deps:
	go mod download

# Package for macOS (creates .app bundle)
package-macos:
	@command -v fyne >/dev/null 2>&1 || (echo "Installing fyne CLI..." && go install fyne.io/fyne/v2/cmd/fyne@latest)
	fyne package -os darwin -name "DataDome Decoder"

# Package for Windows
package-windows:
	@command -v fyne >/dev/null 2>&1 || (echo "Installing fyne CLI..." && go install fyne.io/fyne/v2/cmd/fyne@latest)
	fyne package -os windows -name "DataDome Decoder"

# Package for Linux
package-linux:
	@command -v fyne >/dev/null 2>&1 || (echo "Installing fyne CLI..." && go install fyne.io/fyne/v2/cmd/fyne@latest)
	fyne package -os linux -name "DataDome Decoder"

# Help command
help:
	@echo "Available targets:"
	@echo "  build          - Build the GUI application"
	@echo "  run            - Run the application"
	@echo "  clean          - Remove build artifacts"
	@echo "  install-deps   - Download Go dependencies"
	@echo "  package-macos  - Create macOS .app bundle"
	@echo "  package-windows- Create Windows executable"
	@echo "  package-linux  - Create Linux package"
	@echo "  help           - Show this help message"
