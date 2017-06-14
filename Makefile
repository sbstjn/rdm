run:
	go run *.go

test:
	go test file/*.go

compile:
	go build -ldflags "-X main.version=${VERSION}"

build:
	env GOOS=darwin GOARCH=amd64 make compile

bindata:
	go-bindata -o file/data.go -pkg file templates/*

deploy:
	rm -rf dist	

	mkdir -p dist
	mkdir -p dist/latest

	mv rdm dist/latest/rdm

	zip dist/latest/rdm_darwin_amd64.zip dist/latest/

	shasum -a256 dist/latest/rdm_darwin_amd64.zip

	aws s3 sync dist/latest s3://dl.sbstjn.com/rdm/latest
	aws s3 sync dist/latest s3://dl.sbstjn.com/rdm/${VERSION}