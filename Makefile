GTK_VERSION=$(shell pkg-config --modversion gtk+-3.0 | tr . _ | cut -d '_' -f 1-2)
GTK_BUILD_TAG="gtk_$(GTK_VERSION)"
GIT_VERSION=$(shell git rev-parse HEAD)
TAG_VERSION=$(shell git tag -l --contains $$GIT_VERSION)
GO_VERSION=$(shell go version | grep  -o 'go[[:digit:]]\.[[:digit:]]')

BUILD_DIR=bin

default: gen-ui-defs gen-schema-defs lint test
.PHONY: test

build: build-cli build-gui

# This should not be added as a requirement to build-gui because it may hide
# build problems. build-gui is exactly what `go get` will do on a clean repo
gen-ui-defs:
	make -C ./gui/definitions

gen-schema-defs:
	make -C ./gui/settings/definitions

build-gui: generate-version-file
	go build -i -tags $(GTK_BUILD_TAG) -o $(BUILD_DIR)/coyim

build-gui-win: generate-version-file
	go build -i -tags $(GTK_BUILD_TAG) -ldflags -H=windowsgui -o $(BUILD_DIR)/coyim.exe

build-cli: generate-version-file
	go build -i -tags cli -o $(BUILD_DIR)/coyim-cli

build-debug:
	go build -i -gcflags "-N -l" -tags $(GTK_BUILD_TAG) -o $(BUILD_DIR)/coyim-debug

debug: build-debug
	gdb $(BUILD_DIR)/coyim-debug -d $(shell go env GOROOT) -x build/debug

i18n:
	make -C i18n
.PHONY: i18n

lint:
ifeq ($(GO_VERSION), go1.3)
	echo "Your version of Go is too old for running lint. Skipping."
else
ifeq ($(GO_VERSION), go1.4)
	echo "Your version of Go is too old for running lint. Skipping."
else
	golint ./...
endif
endif

test:
	go test -cover -v -tags $(GTK_BUILD_TAG) ./...

clean-gui-test:
	$(RM) gui-test/*

#TODO: this should only be called on a linux environment
gui-test: clean-gui-test
ifeq ($(shell uname), Linux)
	git clone https://github.com/twstrike/coyim-testing.git gui-test
	echo $$COYIM_PATH
	cd gui-test && behave --stop
endif

generate-version-file:
	./gen_version_file.sh $(GIT_VERSION) $(TAG_VERSION)

run-cover: clean-cover
	mkdir -p .coverprofiles
	go test -coverprofile=.coverprofiles/cli.coverprofile     ./cli
	go test -coverprofile=.coverprofiles/client.coverprofile  ./client
	go test -coverprofile=.coverprofiles/config.coverprofile  ./config
	go test -coverprofile=.coverprofiles/config_importer.coverprofile  ./config/importer
	go test -coverprofile=.coverprofiles/event.coverprofile   ./event
	go test -coverprofile=.coverprofiles/gui.coverprofile  ./gui
	go test -coverprofile=.coverprofiles/i18n.coverprofile    ./i18n
	go test -coverprofile=.coverprofiles/net.coverprofile     ./net
	go test -coverprofile=.coverprofiles/roster.coverprofile  ./roster
	go test -coverprofile=.coverprofiles/sasl.coverprofile    ./sasl
	go test -coverprofile=.coverprofiles/sasl_digestmd5.coverprofile    ./sasl/digestmd5
	go test -coverprofile=.coverprofiles/sasl_plain.coverprofile        ./sasl/plain
	go test -coverprofile=.coverprofiles/sasl_scram.coverprofile        ./sasl/scram
	go test -coverprofile=.coverprofiles/servers.coverprofile ./servers
	go test -coverprofile=.coverprofiles/session.coverprofile ./session
	go test -coverprofile=.coverprofiles/xmpp.coverprofile    ./xmpp
	go test -coverprofile=.coverprofiles/xmpp_data.coverprofile    ./xmpp/data
	go test -coverprofile=.coverprofiles/xmpp_utils.coverprofile    ./xmpp/utils
	go test -coverprofile=.coverprofiles/ui.coverprofile      ./ui
	go test -tags $(GTK_BUILD_TAG) -coverprofile=.coverprofiles/main.coverprofile
	gover .coverprofiles .coverprofiles/gover.coverprofile

clean-cover:
	$(RM) -rf .coverprofiles

# generats an HTML report with coverage information
cover: run-cover
	go tool cover -html=.coverprofiles/gover.coverprofile

get:
	go get -t -tags $(GTK_BUILD_TAG) ./...

deps-dev:
ifeq ($(GO_VERSION), go1.3)
else
ifeq ($(GO_VERSION), go1.4)
else
	go get github.com/golang/lint/golint
endif
endif
	go get golang.org/x/tools/cmd/cover
	go get github.com/modocache/gover
	go get github.com/tools/godep

deps: deps-dev
	godep restore
