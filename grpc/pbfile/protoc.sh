#!/bin/bash

protoc -I .  ./$1 --go_out=plugins=grpc:../services

#protoc -I=. --go_out=. ./pfcp_ie.proto