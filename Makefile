#! /usr/bin/make
#
# Makefile for goa v2
#
# Targets:
# - "depend" retrieves the Go packages needed to run the linter and tests
# - "lint" runs the linter and checks the code format using goimports
# - "test" runs the tests
#
# Meta targets:
# - "all" is the default target, it runs all the targets in the order above.
#
DIRS=$(shell go list -f {{.Dir}} goa.design/goa/expr/...)

# Only list test and build dependencies
# Standard dependencies are installed via go get
DEPEND=\
	github.com/sergi/go-diff/diffmatchpatch \
	golang.org/x/lint/golint \
	golang.org/x/tools/cmd/goimports \
	github.com/golang/protobuf/protoc-gen-go

all: lint gen test

travis: depend all

PROTOC_ZIP = protoc-3.3.0-linux-x86_64.zip
depend:
	# Install protoc
	@curl -s -OL https://github.com/google/protobuf/releases/download/v3.3.0/$(PROTOC_ZIP) && \
		sudo unzip -o $(PROTOC_ZIP) -d /usr/local bin/protoc && \
		rm -f $(PROTOC_ZIP)
	@go get -v $(DEPEND)
	@go install github.com/golang/protobuf/protoc-gen-go
	@go get -t -v ./...

lint:
	@for d in $(DIRS) ; do \
		if [ "`goimports -l $$d/*.go | tee /dev/stderr`" ]; then \
			echo "^ - Repo contains improperly formatted go files" && echo && exit 1; \
		fi \
	done
	@if [ "`golint ./... | grep -vf .golint_exclude | tee /dev/stderr`" ]; then \
		echo "^ - Lint errors!" && echo && exit 1; \
	fi

gen:
	@cd cmd/goa && \
	go install && \
	rm -rf $(GOPATH)/src/goa.design/goa/examples/basic/cmd             && \
	rm -rf $(GOPATH)/src/goa.design/goa/examples/cellar/cmd/cellar-cli && \
	rm -rf $(GOPATH)/src/goa.design/goa/examples/error/cmd             && \
	rm -rf $(GOPATH)/src/goa.design/goa/examples/security/cmd          && \
	rm -rf $(GOPATH)/src/goa.design/goa/examples/streaming/cmd/chatter && \
	goa gen     goa.design/goa/examples/basic/design     -o $(GOPATH)/src/goa.design/goa/examples/basic     && \
	goa example goa.design/goa/examples/basic/design     -o $(GOPATH)/src/goa.design/goa/examples/basic     && \
	goa gen     goa.design/goa/examples/cellar/design    -o $(GOPATH)/src/goa.design/goa/examples/cellar   && \
	goa example goa.design/goa/examples/cellar/design    -o $(GOPATH)/src/goa.design/goa/examples/cellar   && \
	goa gen     goa.design/goa/examples/error/design     -o $(GOPATH)/src/goa.design/goa/examples/error    && \
	goa example goa.design/goa/examples/error/design     -o $(GOPATH)/src/goa.design/goa/examples/error    && \
	goa gen     goa.design/goa/examples/security/design  -o $(GOPATH)/src/goa.design/goa/examples/security && \
	goa example goa.design/goa/examples/security/design  -o $(GOPATH)/src/goa.design/goa/examples/security && \
	goa gen     goa.design/goa/examples/streaming/design -o $(GOPATH)/src/goa.design/goa/examples/streaming  && \
	goa example goa.design/goa/examples/streaming/design -o $(GOPATH)/src/goa.design/goa/examples/streaming

test:
	go test ./...

test-plugins:
	@if [ -z $(GOA_BRANCH) ]; then\
		GOA_BRANCH=$$(git rev-parse --abbrev-ref HEAD); \
	fi
	@if [ ! -d "$(GOPATH)/src/goa.design/plugins" ]; then\
		git clone https://github.com/goadesign/plugins.git $(GOPATH)/src/goa.design/plugins; \
	fi
	@cd $(GOPATH)/src/goa.design/plugins && git checkout $(GOA_BRANCH) || echo "Using master branch in plugins repo" && \
	make -k || (echo "Tests in plugin repo (https://github.com/goadesign/plugins) failed" \
                  "due to changes in goa repo (branch: $(GOA_BRANCH))!" \
                  "Create a branch with name '$(GOA_BRANCH)' in the plugin repo and fix these errors." && exit 1)
