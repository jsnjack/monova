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

dist: build
	@for type in ${BUILD_TYPES} ; do \
		cd ${PWD}/dist && fpm --input-type dir --output-type $$type \
		--name monova --version ${VERSION} --license MIT --no-depends --provides monova \
		--vendor jsnjack@gmail.com \
		--maintainer jsnjack@gmail.com --description \
		"Automatically calculates version of the application based on the commit messages" \
		--url https://github.com/jsnjack/monova --force --chdir ${PWD} ./monova=/usr/bin/monova; \
	done
