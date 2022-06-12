BINARY:=monova
# Check if monova installed
MONOVA:=$(shell monova -version dot 2> /dev/null)
BUILD_TYPES:=rpm deb
PWD:=$(shell pwd)
VERSION=0.0.0

version:
ifdef MONOVA
override VERSION="$(shell monova)"
else
# Use local monova
override VERSION="$(shell test -f monova && ./monova || echo 0.0.0)"
endif

build: ${BINARY} dist/monova_linux_amd64 dist/monova_darwin_amd64

${BINARY}: version
	go build -ldflags="-X main.version=${VERSION}"


dist/monova_linux_amd64: version *.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X main.version=${VERSION}" -o dist/monova_linux_amd64

dist/monova_darwin_amd64: version *.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.version=${VERSION}" -o dist/monova_darwin_amd64

run: build
	./${BINARY}

release: build
	tar --transform='s,_.*,,' --transform='s,dist/,,' -cz -f dist/${BINARY}_linux_amd64.tar.gz dist/${BINARY}_linux_amd64
	tar --transform='s,_.*,,' --transform='s,dist/,,' -cz -f dist/${BINARY}_darwin_amd64.tar.gz dist/${BINARY}_darwin_amd64
	grm release jsnjack/${BINARY} -f dist/${BINARY}_linux_amd64.tar.gz -f dist/${BINARY}_darwin_amd64.tar.gz -t "v`monova`"
