FROM nizsheanez/golang
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

VOLUME ["/gopath"]

RUN mkdir -p /gopath/src/ria

ADD ./ /gopath/src/ria

WORKDIR /gopath/src/ria

RUN go get

RUN go get -v github.com/beego/bee

VOLUME ["/gopath/src/ria"]

EXPOSE ["8080", "80", "8081"]

CMD ["bee", "run"]
