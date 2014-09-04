FROM dockerfile/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

ENV DEBIAN_FRONTEND noninteractive
ENV HOME /root

ADD ./configs /opt/configs/

Run apt-get update
Run apt-get install -y nano zsh

# ZSH
RUN \
        wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O - | zsh && \
        chsh -s `which zsh`

RUN \
        cp /opt/configs/zsh/.zshrc /root/.zshrc && \
        cp /opt/configs/sshd/sshd_config /etc/ssh/sshd_config && \
        rm -rf /opt/configs/

WORKDIR /root

CMD ["zsh"]