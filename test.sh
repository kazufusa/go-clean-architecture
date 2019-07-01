#!/bin/sh
set -Ceux
# go run main.go &
curl -s -w '\n%{http_code}\n' -X GET localhost:8080
curl -s -w '\n%{http_code}\n' -X PUT localhost:8080 -d helloworld
curl -s -w '\n%{http_code}\n' -X GET localhost:8080
