function genProto {
  DOMAIN=$1
  SKIP_GATEWAY=$2
  PROTO_PATH=./${DOMAIN}/api
  GO_OUT_PATH=./${DOMAIN}/api/gen/v1
  mkdir -p $GO_OUT_PATH

  protoc -I=$PROTO_PATH --go_out=paths=source_relative:${GO_OUT_PATH} --go-grpc_out=paths=source_relative:${GO_OUT_PATH} ${DOMAIN}.proto

  if [ $SKIP_GATEWAY ]; then
    return
  fi

  protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=${PROTO_PATH}/${DOMAIN}.yaml:${GO_OUT_PATH} ${DOMAIN}.proto
}

# --swagger_out=allow_repeated_fields_in_body=true,logtostderr=true:${DOMAIN}

genProto auth
genProto rental