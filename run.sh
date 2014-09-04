#!/bin/sh

#Mysql
mysql = docker run --name mysql \
            -p 3306:3306 \
            -v /var/lib/mysql:/var/lib/mysql \
            -d nicescale/percona-mysql

#Golang
app = docker run --name ria \
            -v /gopath:/gopath \
            -p 8080:8080 \
            -p 8081:8081 \
            -p 80:80 \
            --link mysql:db \
            -d nizsheanez/ria

#Bee
app = docker run --name bee \
            --volumes-from ria \
            -d nizsheanez/bee

#Gulp
app = docker run --name gulp \
            --volumes-from ria \
            -d nizsheanez/gulp

echo mysql=$mysql app=$app