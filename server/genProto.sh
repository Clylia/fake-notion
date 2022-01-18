#!/bin/bash
function genProto {
    DOMAIN=$1
    SKIP_GATEWAY=$2
    PROTO_PATH=./${DOMAIN}/api
    GO_OUT_PATH=./${DOMAIN}/api/gen/v1
    mkdir -p $GO_OUT_PATH

    protoc -I $PROTO_PATH --go_out $GO_OUT_PATH --go_opt paths=source_relative --go-grpc_out $GO_OUT_PATH --go-grpc_opt paths=source_relative ${DOMAIN}.proto

    if [ $SKIP_GATEWAY ]; then
        # No external exposure required
        return
    fi

    protoc -I $PROTO_PATH --grpc-gateway_out $GO_OUT_PATH --grpc-gateway_opt paths=source_relative --grpc-gateway_opt grpc_api_configuration=$PROTO_PATH/${DOMAIN}.yaml ${DOMAIN}.proto
}

genProto auth
genProto account
genProto page
genProto blob 1
