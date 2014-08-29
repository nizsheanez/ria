FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

# Install packages
ENV DEBIAN_FRONTEND noninteractive
RUN \
    apt-get update && \
    apt-get -y install mysql-server-5.6 pwgen && \
    apt-get clean && apt-get autoremove && \
    rm -rf /var/lib/apt/lists/*

# Remove pre-installed database
RUN rm -rf /var/lib/mysql/*

# Add MySQL configuration
ADD my.cnf /etc/mysql/conf.d/my.cnf
ADD mysqld_charset.cnf /etc/mysql/conf.d/mysqld_charset.cnf

# Add MySQL scripts
ADD create_mysql_admin_user.sh /create_mysql_admin_user.sh
ADD import_sql.sh /import_sql.sh
ADD run.sh /run.sh
RUN chmod 755 /*.sh

# Exposed ENV
ENV MYSQL_USER root
ENV MYSQL_PASS root

# Add VOLUMEs to allow backup of config and databases
VOLUME  ["/etc/mysql", "/var/lib/mysql"]

EXPOSE ["3306:3306"]
CMD ["/run.sh"]
