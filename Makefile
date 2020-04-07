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

build: version
	go build -ldflags="-X main.version=${VERSION}"

run: build
	./${BINARY}

release: build
	release_on_github -f ${BINARY} -r jsnjack/monova -t "v`monova`"
