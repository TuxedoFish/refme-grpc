# RefMe gRPC

This is a backend gRPC server made for use within the RefMe app. 

# Setup

To get started:

```git clone --recurse-submodules -j8 git@github.com:TuxedoFish/refme-grpc.git```

Then fill in the details within the .env file

Then to run the server locally:

```make run```

# gRPC services

## GetArticles

Takes ArticlesPageRequest:

```
message ArticlesPageRequest {
    string query_string = 1;
    optional int32 page = 2;
}
```

and returns an ArticlesPageResponse:

```
message Meta {
    repeated Provider providers = 1;
    string query = 2;
    int32 page = 3;
    int32 results = 4;
}

message Result {
    string id = 1;
    string author = 2;
    string title = 3;
    string published_date = 4;
    string publisher = 5;
    string description = 6;
    string url = 7;
}

message ArticlesPageResponse {
    Meta meta = 1;
    repeated Result results = 2;
}
```

This is done by making requests to 3 APIs using the D'Hondt splitting system to decide based on weightings how to split the request up. The 3 APIs are x5gon, arXiv and springer. arXiv is an example of an API that returns XML which requires unmarshalling in order to use. Both x5gon and springer send back to json the only difference being springer requires an API key.

# Running tests

To run the tests:

```gotestsum --format testname```

Or to run an individual test suite:

```gotestsum --format testname ./internal/articles```

# Updating services

In the event that it is needed to update the services due to a protobuffer change:

```make build```