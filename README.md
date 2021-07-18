# RefMe gRPC

This is a backend gRPC server made for use within the RefMe app. 

# Setup

To get started:

```git clone --recurse-submodules -j8 git@github.com:TuxedoFish/refme-grpc.git```

Then to run the server locally:

```go run service.go```

# Updating services

In the event that it is needed to update the services due to a protobuffer change:

```./generate.sh```