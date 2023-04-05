#!/usr/bin/env bash

export GOBIN=$PWD/pkg/config/tools
export PATH=$GOBIN:$PATH

protoc --go_out=. --twirp_out=. ./contracts/channel/channel.proto  --experimental_allow_proto3_optional
protoc --go_out=. --twirp_out=. ./contracts/participant/participant.proto  --experimental_allow_proto3_optional
protoc --go_out=. ./contracts/wsMessages/models.proto  --experimental_allow_proto3_optional
protoc --go_out=. ./contracts/wsMessages/wsMessages.proto -I "./contracts/wsMessages" --experimental_allow_proto3_optional