FROM dockerfile/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

ENV DEBIAN_FRONTEND noninteractive

# Basic soft
RUN  \
        apt-get update && apt-get upgrade -y && apt-get dist-upgrade -y && \
        apt-get clean -y && apt-get autoremove -y

Run \
        apt-get update && \
        apt-get install -y nano tmux openssh-server zsh && \
        apt-get clean && apt-get autoremove

ADD ./configs /opt/configs/

# ZSH
RUN \
        wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O - | zsh && \
        chsh -s `which zsh`

RUN \
        cp /opt/configs/zsh/.zshrc /root/.zshrc && \
        cp /opt/configs/sshd/sshd_config /etc/ssh/sshd_config


ENTRYPOINT ["/bin/zsh"]
