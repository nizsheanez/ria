FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

# SSH - we need it because no NFS support in go. See https://github.com/howeyc/fsnotify/issues/107
RUN \
        apt-get install -y --no-install-recommends openssh-server

ADD ./configs /opt/configs/

RUN \
     cp /opt/configs/sshd/sshd_config /etc/ssh/sshd_config && \
    rm -rf /opt/configs/

EXPOSE 22

CMD ["zsh"]