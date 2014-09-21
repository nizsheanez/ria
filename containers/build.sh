#!/bin/bash

docker build -t nizsheanez/ubuntu ./ubuntu

docker build -t nizsheanez/mysql ./mysql &
docker build -t nizsheanez/golang ./golang  &
docker build -t nizsheanez/gulp ./gulp  &
docker build -t nizsheanez/bower ./bower &
docker build -t nizsheanez/debug ./debug &
docker build -t nizsheanez/zsh ./zsh &

cd ../ && docker build -t nizsheanez/ria .


docker rm $(docker ps -a -q)

docker rmi $(docker images -q --filter "dangling=true")