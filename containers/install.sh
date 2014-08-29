#!/bin/sh
docker run --name mysql -p 3306:3306 -v /var/lib/mysql:/var/lib/mysql -d nizsheanez/mysql