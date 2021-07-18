# RefMe gRPC

This is a backend gRPC server made for use within the RefMe app. 

# Setup

To get started:

```git clone --recurse-submodules -j8 git@github.com:TuxedoFish/refme-grpc.git```

Then fill in the details within the .env file

Then to run the server locally:

```make run```

# Running tests

To run the tests:

```gotestsum --format testname```

# Updating services

In the event that it is needed to update the services due to a protobuffer change:

```make build```