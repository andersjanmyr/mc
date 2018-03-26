#!/bin/bash -e

ORG_PATH="github.com/andersjanmyr"
REPO_PATH="${ORG_PATH}/mc"

export GOPATH=${PWD}/gopath

rm -f $GOPATH/src/${REPO_PATH}
mkdir -p $GOPATH/src/${ORG_PATH}
ln -s ${PWD} $GOPATH/src/${REPO_PATH}

eval $(go env)
go get \
  github.com/bradfitz/gomemcache/memcache \
  github.com/mitchellh/go-homedir \
  github.com/pkg/errors \
  github.com/spf13/cobra \
  github.com/spf13/viper

CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o bin/mc ${REPO_PATH}
