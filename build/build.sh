#!/usr/bin/env bash
echo compiling Go via Docker
docker run -it --rm -v ${GOPATH}:/app -e GOPATH="/app" -w "/app/src/github.com/nethousejela/hellogo" golang sh -c 'CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o build/hello'
echo Building Docker image

docker build -t nethousejela/hellogo .
docker push nethousejela/hellogo

#set HOST to "DB", PORT to 5984, link "DB" to CouchDB container
#docker run -p 8080:8080 -e HOST=DB -e PORT=5984 --link 991dfe34173f:DB nethousejela/hellog