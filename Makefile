BINARY_NAME=godoro

LINUX_SUF=_linux
MAC_SUF=_macos
WIN_SUF=_windows.exe

run:
	@go run main.go

build: clean build-linux build-macos build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)$(LINUX_SUF)

build-macos:
	GOOS=darwin GOARCH=arm64 go build -o ./bin/$(BINARY_NAME)$(MAC_SUF)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)$(WIN_SUF)

clean:
	rm -f $(BINARY_NAME)$(LINUX_SUF)
	rm -f $(BINARY_NAME)$(MAC_SUF)
	rm -f $(BINARY_NAME)$(WIN_SUF)

.PHONY: build build-linux build-macos build-windows run clean
