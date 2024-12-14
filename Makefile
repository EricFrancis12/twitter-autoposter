BIN_FILE_PATH := ./twitter-autoposter

build:
	go build -o $(BIN_FILE_PATH)

# Example pass in arguments:
# make run ARGS="-mint=30 -maxt=60 -errt=40"
run: build
	$(BIN_FILE_PATH) $(ARGS)

test:
	go test -v ./...

create_config:
	./scripts/create_config.sh
