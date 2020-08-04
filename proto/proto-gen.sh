#!/bin/sh

if test -z "$1"
then
  PROTO_FILES=$(find ./ -name "*.proto" -not -path "*/vendor/*")
else
  PROTO_FILES=$(find "$1" -name "*.proto" -not -path "*/vendor/*")
fi

for PROTO_FILE in ${PROTO_FILES}; do
    echo "protoc -> $PROTO_FILE"
    protoc -I=/usr/local/include -I. \
        --go_out=. \
        "${PROTO_FILE}"
done

