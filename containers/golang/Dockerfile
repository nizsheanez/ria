FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

RUN apt-get update

RUN \
    apt-get install -y golang-go

RUN \
    git clone https://github.com/pote/gpm.git /tmp/gpm && cd /tmp/gpm && \
    git checkout v1.2.3 && ./configure && make install && \
    rm -rf /tmp/gpm && \
    \
    git clone https://github.com/pote/gvp.git /tmp/gvp &&  cd /tmp/gvp && \
    git checkout tags/0.1.0 && ./configure && make install && \
    rm -rf /tmp/gpm-git


RUN \
    mkdir /goroot /gopath

ENV GOROOT /goroot
ENV GOPATH /gopath

ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin