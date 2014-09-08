FROM nizsheanez/golang
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

RUN go get -v github.com/revel/cmd/revel

WORKDIR /gopath/src/ria2

EXPOSE 9000

CMD ["revel", "run", "ria2"]
