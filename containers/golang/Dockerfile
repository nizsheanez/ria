FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

RUN apt-get update

RUN \
    apt-get install -y golang-go

RUN \
    mkdir /goroot /gopath

ENV GOROOT /goroot
ENV GOPATH /gopath

ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin