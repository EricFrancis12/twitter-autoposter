BIN_FILE_PATH := ./twitter-autoposter

build:
	go build -o $(BIN_FILE_PATH)

run: build
	$(BIN_FILE_PATH)

test:
	go test -v ./...

create_config:
	./scripts/create_config.sh
