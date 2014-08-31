#!/bin/sh
vagrant resume

#Mysql
mysql = docker run --name mysql \
            -p 3306:3306 \
            -v /var/lib/mysql:/var/lib/mysql \
            -v /etc/mysql:/etc/mysql \
            -d nizsheanez/mysql

#Golang
app = docker run --name mysql \
            -v /projects/ria:/gopath/src/ria \
            -d nizsheanez/ria

echo mysql=$mysql app=$app