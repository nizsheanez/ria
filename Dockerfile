FROM google/golang
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

VOLUME ["/gopath/src/ria"]

EXPOSE ["8080", "80", "8081"]

WORKDIR /gopath/src/ria
