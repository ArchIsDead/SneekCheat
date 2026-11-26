BIN = sneekcheat
PREFIX ?= /data/data/com.termux/files/usr

all: build

build:
	go build -ldflags="-s -w" -o $(BIN) .

install: build
	install -Dm755 $(BIN) $(PREFIX)/bin/$(BIN)

deps:
	pkg install -y golang termux-api

run: build
	./$(BIN)

clean:
	rm -f $(BIN)

.PHONY: all build install deps run clean
