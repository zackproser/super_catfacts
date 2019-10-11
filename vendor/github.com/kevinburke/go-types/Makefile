.PHONY: install test
BUMP_VERSION := $(GOPATH)/bin/bump_version
GODOCDOC := $(GOPATH)/bin/godocdoc
MEGACHECK := $(GOPATH)/bin/megacheck

test: lint
	go test ./...

install:
	go get ./...
	go install ./...

$(MEGACHECK):
	go get honnef.co/go/tools/cmd/megacheck

lint: | $(MEGACHECK)
	$(MEGACHECK) ./...
	go vet ./...

race-test:
	go test -race ./...

$(BUMP_VERSION):
	go get github.com/kevinburke/bump_version

release: test | $(BUMP_VERSION)
	$(BUMP_VERSION) minor types.go

docs: | $(GODOCDOC)
	$(GODOCDOC)
