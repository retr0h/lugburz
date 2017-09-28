VENDOR := vendor
BINDATA := bindata/go-bindata
GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
GITCOMMIT := $(GITCOMMIT)-dirty
endif
VERSION := $(shell cat VERSION)
LDFLAGS := \
	-s \
	-w \
	-X main.GITCOMMIT=${GITCOMMIT} \
	-X main.VERSION=${VERSION} \

test: go-bindata-dev fmt lint vet
	@echo "+ $@"
	go test -covermode=count ./...

fmt:
	@echo "+ $@"
	@gofmt -s -l . | grep -v $(VENDOR) | tee /dev/stderr

lint:
	@echo "+ $@"
	@golint ./... | grep -v $(VENDOR) | grep -v resource/asset.go | tee /dev/stderr

vet:
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v $(VENDOR))

clean:
	@echo "+ $@"
	@rm -rf ./build
	@rm -f ${BINDATA}

build: clean go-bindata
	@echo "+ $@"
	gox \
		-osarch="linux/amd64 darwin/amd64" \
		-ldflags="${LDFLAGS}" \
		-output="build/{{.Dir}}_{{.OS}}_{{.Arch}}"

${BINDATA}:
	go build -o $@ ./vendor/github.com/jteeuwen/go-bindata/go-bindata/...

.PHONY: build-bindata
build-bindata: ${BINDATA}

go-bindata: ${BINDATA}
	@echo "+ $@"
	${BINDATA} -pkg resource -o resource/asset.go resource/resource_schema_v1.json

go-bindata-dev: ${BINDATA}
	@echo "+ $@"
	${BINDATA} -debug -pkg resource -o resource/asset.go resource/resource_schema_v1.json
