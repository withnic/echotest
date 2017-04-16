# メタ情報
NAME := echotest
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
	-X 'main.revision=$(REVISION)'

# 必要なパッケージのインストール
### Setup
setup:
	go get github.com/golang/dep
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/stretchr/testify/assert
	go get github.com/labstack/echo
	go get github.com/mattn/go-sqlite3
	go get github.com/go-gorp/gorp
	go get bitbucket.org/liamstask/goose/cmd/goose
	go get github.com/go-playground/validator

# テスト実行
test: setup
	go test $$(go list ./...|grep -v vendor)

# depsを使って依存をインストール
deps: setup
	dep ensure

## Update dependencies
update: setup
	dep ensure -update

## lint
##lint: setup
##	go vet $$(go list ./...|grep -v vendor)
##	for pkg in $$(glide novendor -x); do \
##		golint -set_exit_status $$pkg || exit $$?; \
##	done

## Format source codes
##fmt: setup
##	goimports -w $$(glide nv -x)

## buid binaries ex, make bin/echotest
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<

