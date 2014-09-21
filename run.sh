#!/bin/sh
#Mysql
mysql=docker run --name mysql \
            -p 3306:3306 \
            -v /var/lib/mysql:/var/lib/mysql \
            -d nizsheanez/mysql

#Golang
app=docker run --name ria \
            -p 8080:8080 \
            -p 8081:8081 \
            -p 80:80 \
            -p 22:22 \
            --link mysql:db \
            -i nizsheanez/ria

docker run --rm \
            --vaues-from ria \
            -ti nizsheanez/bower bash -c 'cd /gopath/src/ria && bower --allow-root install'

docker run --rm \
            --volumes-from ria \
            -i nizsheanez/bower 

#RUN cd /gopath/src/ria/static && bower install --allow-root
#Bee
bee=docker run --name bee \
            --volumes-from ria \
            -d nizsheanez/bee

#Gulp
gulp=docker run --name gulp \
            --volumes-from ria \
            -d nizsheanez/gulp

echo mysql=$mysql app=$app bee=$bee gulp=$gulp