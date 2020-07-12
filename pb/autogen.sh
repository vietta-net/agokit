#! /bin/bash

cd ..
cd ..
protoc --go_out=plugins=grpc:./ ./agokit/pb/*.proto
