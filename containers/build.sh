#!/bin/bash

docker pull dockerfile/ubuntu &
docker pull dockerfile/elasticsearch &
docker pull dockerfile/mysql &
docker pull dockerfile/percona &
docker pull dockerfile/nginx &
docker pull dockerfile/nodejs-bower-gulp &

wait

docker build -t nizsheanez/ubuntu ./ubuntu

docker build -t nizsheanez/golang ./golang  &
docker build -t nizsheanez/mysql ./mysql &
docker build -t nizsheanez/nodejs ./nodejs &

wait

docker build -t nizsheanez/gulp ./gulp  &
docker build -t nizsheanez/bower ./bower &
docker build -t nizsheanez/debug ./debug &

wait

cd ../ && docker build -t nizsheanez/ria .


docker rm $(docker ps -a -q)

docker rmi $(docker images -q --filter "dangling=true")