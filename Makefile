MAKEFLAGS += --silent

GO_BINARY= service


# LOCAL DEV
clean:
	rm -rf $(GO_BINARY)
	
build: clean
	cd cmd/service \
	&& go build -o ../../$(GO_BINARY)

run: build
	./$(GO_BINARY)



# TODO: set up docker commands