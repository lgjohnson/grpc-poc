#!/usr/bin/env bash

#create go code from proto
protoc \
    -I echolalia/ \
    echolalia/echolalia.proto \
    --go_out=plugins=grpc:echolalia

#copy go code to GOPATH
mkdir -p $GOPATH/src/github.com/gjohnson/echolalia
cp echolalia/echolalia.pb.go $GOPATH/src/github.com/gjohnson/echolalia

#copy server code to GOPATH
mkdir -p $GOPATH/src/github.com/gjohnson/server
cp server/main.go $GOPATH/src/github.com/gjohnson/server/

# go install server
go install $GOPATH/src/github.com/gjohnson/server


# run server
$GOPATH/bin/server 

# run client



#cp echolalia/echolalia.pb.go $GOPATH/src/github.com/gjohnson/echolalia

