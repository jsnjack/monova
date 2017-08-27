BINARY:=monova
MONOVA:=$(shell dahary -v dot 2> /dev/null)
VERSION:=0.0.0

version:
ifdef MONOVA
	VERSION:=$(shell monova)
else
	$(info !Install monova to automatically update version)
endif

build: version
	go build -ldflags="-X main.version=${VERSION}"

run: build
	./${BINARY}
