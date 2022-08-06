#!/bin/sh
source_path=./cmd/build
go_file=main.go
image_name=yinwan
build_output=yiwan
version=1.0.1


docker rmi -f $image_name:$version
docker build -f $source_path/Dockerfile -t $image_name:$version  .
rm  $source_path/$build_output

docker save -o $image_name-$version.tar $image_name:$version
