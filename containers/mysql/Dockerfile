FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

# Install packages
ENV DEBIAN_FRONTEND noninteractive
RUN \
        apt-get update && \
        apt-get -y install mysql-server-5.6 pwgen && \
        apt-get clean && apt-get autoremove && \
        rm -rf /var/lib/apt/lists/*

ADD ./configs /opt/configs/

# Add MySQL configuration
# Add MySQL scripts
# We do rm /var/lib/mysql* after mysql installation, because will use data from Volume
RUN \
        mv /opt/configs/cnf/my.cnf /etc/mysql/conf.d/my.cnf && \
        mv /opt/configs/cnf/mysqld_charset.cnf /etc/mysql/conf.d/mysqld_charset.cnf && \
        mv /opt/configs/scripts/create_mysql_admin_user.sh /create_mysql_admin_user.sh && \
        mv /opt/configs/scripts/import_sql.sh /import_sql.sh && \
        mv /opt/configs/scripts/run.sh /run.sh && \
        rm -rf /opt/configs && \
        chmod 755 /*.sh && \
        rm -rf /var/lib/mysql/*

# Exposed ENV
ENV MYSQL_USER root
ENV MYSQL_PASS root

# Add VOLUMEs to allow backup of config and databases
VOLUME  ["/etc/mysql", "/var/lib/mysql"]

EXPOSE 3306
CMD ["/run.sh"]
