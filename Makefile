ROOT := $(CURDIR)
GOPKGS = \
		golang.org/x/tools/cmd/cover \
		github.com/golang/lint/golint \
		github.com/tools/godep

default: test

deps:
	@go get -u -v $(GOPKGS)
	@if [ `which godep` ] && [ -f ./Godeps/Godeps.json ]; then godep restore; fi

lint:
	@echo "[Lint] running golint"
	@golint

vet:
	@echo "[Vet] running go vet"
	@go vet

ci: deps vet lint test

test:
	@echo "[Test] running tests"
	@go test -v ./... -cover -coverprofile=c.out

.PHONY: default golint test
