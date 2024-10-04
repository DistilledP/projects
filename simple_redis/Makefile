run-server:
	@LUNGFISH_PUBLIC_PORT=6379 go run cmd/server/main.go

build-server:
	@go build -o build/server/simple_redis cmd/server/main.go

clean:
	@rm -fr build

fix-imports: # install with:  go install -v github.com/incu6us/goimports-reviser/v3@latest
	@goimports-reviser -excludes vendor/ -recursive ./
