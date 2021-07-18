#!/bin/bash

protoc refme-protobuf/articles/articles.proto --go_out=plugins=grpc:.