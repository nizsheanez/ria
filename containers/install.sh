#!/bin/sh

#Mysql
docker run --name mysql \
        -p 3306:3306 \
        -v /var/lib/mysql:/var/lib/mysql \
        -v /etc/mysql:/etc/mysql \
        -d nizsheanez/mysql

#Golang
docker run --name mysql \
        -v /var/lib/mysql:/var/lib/mysql \
        -v /etc/mysql:/etc/mysql \
        -d nizsheanez/mysql
