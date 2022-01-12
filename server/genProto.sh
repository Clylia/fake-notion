#!/bin/bash
function genPeoto {
    DOMAIN=$1
    SKIP_GATEWAY=$2
    PROTO_PATH=./${DOMAIN}/api
    GO_OUT_PATH=./${DOMAIN}/api/gen/v1
    mkdir -p $GO_OUT_PATH

    protoc -I $PROTO_PATH --go_out $GO_OUT_PATH --go_opt paths=source_relative --go-grpc_out $GO_OUT_PATH --go-grpc_opt paths=source_relative ${DOMAIN}.proto
}

genPeoto auth
genPeoto account
