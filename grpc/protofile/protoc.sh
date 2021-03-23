#!/bin/bash
DIR=`pwd`
#echo $DIR/$1
protoc  --go_out=plugins=grpc:../pbfile -I $DIR/$1
#protoc -I ${GOPATH}/src  --go_out=plugins=grpc:../pbfile  $DIR/$1
#protoc -I .  ./$1 --go_out=plugins=grpc:../pbfile
