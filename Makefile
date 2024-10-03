default: ; go build -o build/ .
windows: ; GOOS=windows GOARCH=amd64 go build -o build/ .

clean: ; rm -rf build/
