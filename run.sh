#!/bin/sh
#Mysql
mysql=docker run --name mysql \
            -p 3306:3306 \
            -v /var/lib/mysql:/var/lib/mysql \
            -i dockerfile/mysql

#Golang
app=docker run --name ria \
            -p 8080:8080 \
            -p 8081:8081 \
            -p 80:80 \
            -p 22:22 \
            --link mysql:db \
            -i nizsheanez/ria


app=docker run --name ria2 \
            -v /gopath/src/ria2:/gopath/src/ria2 \
            -d nizsheanez/ria2

docker run --name revel \
            --volumes-from ria2 \
            -p 9000:9000 \
            -i nizsheanez/revel

#Bee
bee=docker run --name bee \
            --volumes-from ria \
            -d nizsheanez/bee

#Gulp
gulp=docker run --name gulp \
            --volumes-from ria \
            -d nizsheanez/gulp

echo mysql=$mysql app=$app bee=$bee gulp=$gulp