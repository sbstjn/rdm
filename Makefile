run:
	go run *.go

test:
	go test -cover -race ./...

bench:
	go test -bench=. ./...

race:
	go test -v -race ./...

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

	cd dist/latest && zip rdm_darwin_amd64.zip rdm

	shasum -a256 dist/latest/rdm_darwin_amd64.zip

	aws s3 sync dist/latest s3://dl.sbstjn.com/rdm/latest
	aws s3 sync dist/latest s3://dl.sbstjn.com/rdm/${VERSION}