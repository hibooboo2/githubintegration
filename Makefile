BIN=githubintegrations
BUILD := $(shell [[ `git log -n1 --pretty='%h' | xargs git describe --exact-match --tags` != "" ]] && git describe --tags || git rev-parse --verify HEAD)
SERVICE=hibooboo2-${BIN}
IMAGE=hibooboo2/${BIN}

default: build-local

deps:
	echo Need to use

build-linux: deps
	CGO_ENABLED=0 GOOS=linux go build -o ${BIN} -a -tags netgo -ldflags '-s -w -X main.Build="${BUILD}"'

build-local: deps
	CGO_ENABLED=0 go build -o ${BIN} -a -tags netgo -ldflags '-s -w -X main.Build="${BUILD}"'

version:
	echo Build is: ${BUILD}

linode: build-linux
	scp ${BIN} linode.wizardofmath.host:~
