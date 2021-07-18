.PHONY: compile
compile: ## Compile the proto file.
	./generate.sh
 
.PHONY: run
run: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server cmd/server/main.go
	bin/server