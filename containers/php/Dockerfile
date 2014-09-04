FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

# Basic soft
RUN  \
        apt-get install -y php5 php5-fpm php5-dev && \
        apt-get install -y php5-xcache php5-gd php5-imagick php5-mysqlnd && \
        apt-get install -y php5-xdebug php5-curl php5-mcrypt

ADD ./configs /opt/configs/

# ZSH
RUN \
        wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O - | zsh && \
        chsh -s `which zsh`

RUN \
        cp /opt/configs/zsh/.zshrc /root/.zshrc && \
        cp /opt/configs/sshd/sshd_config /etc/ssh/sshd_config


ENTRYPOINT ["/bin/zsh"]
