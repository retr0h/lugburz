clean:
	@rm -rf ./build

build: clean
	@$(GOPATH)/bin/goxc \
	  -bc="darwin,amd64" \
	  -d=build
