ARGS1 := $(word 1,$(MAKECMDGOALS))
ARGS2 := $(word 2,$(MAKECMDGOALS))

TARGETOS ?= $(if $(filter apple,$(ARGS1)),darwin,$(if $(filter windows,$(ARGS1)),windows,linux))
TARGETARCH ?= $(if $(filter arm arm64,$(ARGS2)),arm64,$(if $(filter amd amd64,$(ARGS2)),amd64,amd64))

BUILD_DIR ?= build
EXT ?= 

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get:
	go get

build: format get ## Default build for Linux amd64
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o ${BUILD_DIR}/telebot_${TARGETOS}${EXT} -ldflags "-X=github.com/andriuk/telebot/cmd.appVersion=${VERSION}"

linux: build ## Build a Linux binary. [ linux [[arm|arm64] | [amd|amd64]] ] to build for the specific ARCH 

apple: build ## Build a macOS binary

windows: EXT = .exe
windows: build ## Build a Windows binary

%::
	@true
