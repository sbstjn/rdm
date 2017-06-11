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
	cp rdm dist/latest/rdm
	cd dist/latest && zip rdm_darwin_amd64.zip rdm
	aws s3 sync dist/latest s3://dl.sbstjn.com/rdm/latest
	mv dist/latest dist/${VERSION}
	aws s3 sync dist/${VERSION} s3://dl.sbstjn.com/rdm/${VERSION}