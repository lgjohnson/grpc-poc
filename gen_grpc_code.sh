#!/usr/bin/env bash

protoc \
    -I echolalia/ \
    echolalia.proto \
    --go_out=plugins=grpc:echolalia
