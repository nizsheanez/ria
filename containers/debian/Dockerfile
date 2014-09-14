FROM debian:wheezy
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

ENV DEBIAN_FRONTEND noninteractive

# Install.
RUN \
  sed -i 's/# \(.*multiverse$\)/\1/g' /etc/apt/sources.list && \
  apt-get update && \
  apt-get -y upgrade && \
  apt-get install -y build-essential && \
  apt-get install -y software-properties-common && \
  apt-get install -y byobu curl git htop man unzip vim wget


# Set environment variables.
ENV HOME /root

# Define working directory.
WORKDIR /root


Run \
        apt-get install -y --no-install-recommends nano openssh-server zsh

ADD ./configs /opt/configs/

# ZSH
RUN \
        wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O - | zsh && \
        chsh -s `which zsh`

RUN \
        mkdir /var/run/sshd && \
        cp /opt/configs/sshd/sshd_config /etc/ssh/sshd_config && \
        cp /opt/configs/zsh/.zshrc /root/.zshrc

EXPOSE 22

ENTRYPOINT ["/bin/zsh"]
