#!/bin/sh
docker run -d \
    -v /var/lib/mysql:/var/lib/mysql \
    -e MYSQL_PASS="asharov" \
    tutum/mysql

docker run --rm \
            -v /var/www/ria:/gopath/src/ria \
            -ti nizsheanez/golang bash -c 'cd /gopath/src/ria && gvp init && source gvp in && gpm install'

docker run --rm \
            -v /var/www/ria:/gopath/src/ria \
            -ti nizsheanez/bower bash -c 'cd /gopath/src/ria/static && bower --allow-root install'

#Mysql
mysql=docker run --rm --name mysql \
            -p 3306:3306 \
            -v /var/lib/mysql:/var/lib/mysql \
            -d tutum/mysql

#Golang
app=docker run --rm --name ria \
            -p 8080:8080 \
            -p 8081:8081 \
            -p 80:80 \
            -v /var/www/ria:/gopath/src/ria \
            -i nizsheanez/ria

#Gulp
gulp=docker run --name gulp \
            --volumes-from ria \
            -d nizsheanez/gulp

echo mysql=$mysql app=$app bee=$bee gulp=$gulp