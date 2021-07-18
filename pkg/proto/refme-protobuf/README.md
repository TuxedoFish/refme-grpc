# RefMe Protobuffers

This is a repository containing all of the protobuffer files for the RefMe project 

# Guide

* All protobuffers use syntax="proto3". 

* In order to validate that protobuffers have correct syntax please run:

```
protoc --lint_out=. */*.proto
```