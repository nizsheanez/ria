FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

# Install Percona Server.
RUN \
  apt-key adv --keyserver keys.gnupg.net --recv-keys 1C4CBDCDCD2EFD2A && \
  echo "deb http://repo.percona.com/apt `lsb_release -cs` main" > /etc/apt/sources.list.d/percona.list && \
  apt-get update

RUN \
  apt-get install -y percona-server-server-5.6

ADD ./configs /opt/configs/mysql/

RUN \
  chmod 0755 /opt/configs/mysql/scripts -R && \
#  rm -rf /var/lib/apt/lists/* && \
  cp /opt/configs/mysql/cnf/my.cnf /etc/mysql/my.cnf && \
  echo "mysqld_safe &" > /tmp/config && \
  echo "mysqladmin --silent --wait=30 ping || exit 1" >> /tmp/config && \
  echo "mysql -e 'GRANT ALL PRIVILEGES ON *.* TO \"root\"@\"%\" WITH GRANT OPTION;'" >> /tmp/config && \
  echo "mysql -e 'create user \"root\"@\"10.0.2.2\" identified by \"yourpassword\";'" >> /tmp/config && \
  echo "mysql -e 'grant all privileges on *.* to \"root\"@\"10.0.2.2\" with grant option;'" >> /tmp/config && \
  echo "mysql -e 'flush privileges;'" >> /tmp/config && \
  bash /tmp/config && \
  rm -f /tmp/config

# Define mountable directories.
VOLUME ["/etc/mysql", "/var/lib/mysql"]

EXPOSE 3306

WORKDIR /data

ENTRYPOINT ["/opt/configs/mysql/scripts/run.sh"]

CMD ["mysqld", "--datadir=/var/lib/mysql", "--user=mysql"]