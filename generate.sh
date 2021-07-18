#!/bin/bash

protoc pkg/proto/refme-protobuf/articles/articles.proto --go_out=plugins=grpc:pkg/proto/