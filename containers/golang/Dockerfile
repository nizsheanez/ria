FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

RUN \
    apt-get update && \
    apt-get -y install golang && \
    apt-get clean

RUN \
    mkdir /goroot

ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin