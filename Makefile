PKGS ?= $(shell go list ./... | /usr/bin/grep -v /vendor/)
PKGS_DELIM ?= $(shell echo $(PKGS) | sed -e 's/ /,/g')
VENDOR := vendor
BINDATA := go-bindata
GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
GITCOMMIT := $(GITCOMMIT)-dirty
endif
VERSION := $(shell cat VERSION)
LDFLAGS := \ -s \
	-w \
	-X main.GITCOMMIT=${GITCOMMIT} \
	-X main.VERSION=${VERSION} \

test: go-bindata-dev fmt lint vet
	@echo "+ $@"
	go test -covermode=count ./...

cover:
	@echo "+ $@"
	$(shell [ -e cover.out ] && rm cover.out)
	@go list -f '{{if or (len .TestGoFiles) (len .XTestGoFiles)}}go test -test.v -test.timeout=120s -covermode=count -coverprofile={{.Name}}_{{len .Imports}}_{{len .Deps}}.coverprofile -coverpkg $(PKGS_DELIM) {{.ImportPath}}{{end}}' $(PKGS) | xargs -I {} bash -c {}
	@echo "mode: count" > cover.out
	@grep -h -v "^mode:" *.coverprofile >> "cover.out"
	@rm *.coverprofile
	@go tool cover -html=cover.out -o=cover.html

fmt:
	@echo "+ $@"
	@gofmt -s -l . | grep -v $(VENDOR) | grep -v resource/asset.go | tee /dev/stderr

lint:
	@echo "+ $@"
	@golint ./... | grep -v $(VENDOR) | grep -v resource/asset.go | tee /dev/stderr

vet:
	@echo "+ $@"
	@go vet $(shell go list ./... | grep -v $(VENDOR))

clean:
	@echo "+ $@"
	@rm -rf ./build

build: clean go-bindata
	@echo "+ $@"
	gox \
		-osarch="linux/amd64 darwin/amd64" \
		-ldflags="${LDFLAGS}" \
		-output="build/{{.Dir}}_{{.OS}}_{{.Arch}}"

go-bindata:
	@echo "+ $@"
	$(BINDATA) -pkg resource -o resource/asset.go resource/resource_schema_v1.json

go-bindata-dev:
	@echo "+ $@"
	$(BINDATA) -debug -pkg resource -o resource/asset.go resource/resource_schema_v1.json
