FROM google/golang:latest
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

RUN go get -v github.com/beego/bee

WORKDIR /gopath/src/ria

CMD ["bee", "run"]
