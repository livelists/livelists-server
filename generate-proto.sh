#!/usr/bin/env bash

export GOBIN=$PWD/pkg/config/tools
export PATH=$GOBIN:$PATH

protoc --go_out=. --twirp_out=. ./contracts/channel/channel.proto