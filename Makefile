run-server:
	@LUNGFISH_PUBLIC_PORT=6379 go run cmd/server/main.go

build-server:
	@go build -o build/server/lungfish cmd/server/main.go

cli: 
	@go run cmd/cli/main.go

build-cli:
	@go build -o build/cli/lungfish-cli cmd/cli/main.go

clean:
	@rm -fr build

fix-imports: # install with:  go install -v github.com/incu6us/goimports-reviser/v3@latest
	@goimports-reviser -excludes vendor/ -recursive ./
