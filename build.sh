#!/bin/bash
if [ $GOPATH == ""] 
then
    GOPATH=$HOME/go
fi
PATH=$GOPATH/bin/:$PATH
go get github.com/jteeuwen/go-bindata/...

go-bindata ./frontend/...
go build ./