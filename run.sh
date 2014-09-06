#!/bin/sh

#Mysql
mysql=docker run --name mysql \
            -p 3306:3306 \
            -v /var/lib/mysql:/var/lib/mysql \
            -d nizsheanez/percona

#Golang
app=docker run --name ria \
            -v /gopath:/gopath \
            -p 8080:8080 \
            -p 8081:8081 \
            -p 80:80 \
            --link mysql:db \
            -d nizsheanez/ria

#Bee
bee=docker run --name bee \
            --volumes-from ria \
            -d nizsheanez/bee

#Gulp
gulp=docker run --name gulp \
            --volumes-from ria \
            -d nizsheanez/gulp

echo mysql=$mysql app=$app bee=$bee gulp=$gulp