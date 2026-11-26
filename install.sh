#!/data/data/com.termux/files/usr/bin/sh
set -e
pkg install -y golang termux-api
go build -ldflags="-s -w" -o "$PREFIX/bin/sneekcheat" .
echo "done, run: sneekcheat"
