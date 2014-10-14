FROM nizsheanez/golang
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

EXPOSE 8080
EXPOSE 80
EXPOSE 8081

VOLUME ["/gopath/src/ria"]

WORKDIR /gopath/src/ria

CMD ["go build && ls -la"]

