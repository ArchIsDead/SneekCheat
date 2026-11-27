BIN = sneekcheat
PREFIX ?= /data/data/com.termux/files/usr

all: build

build:
	go build -ldflags="-s -w" -o $(BIN) .

vet:
	go vet ./...

install: build
	install -Dm755 $(BIN) $(PREFIX)/bin/$(BIN)

deps:
	pkg install -y golang termux-api

run: build
	./$(BIN)

belakang: build
	./$(BIN) jalan

mati: build
	./$(BIN) mati

clean:
	rm -f $(BIN)

.PHONY: all build vet install deps run belakang mati clean
