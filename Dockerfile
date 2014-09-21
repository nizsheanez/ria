FROM nizsheanez/golang
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

ADD ./ /gopath/src/ria

RUN cd /gopath/src/ria && gpm install
#RUN cd /gopath/src/ria/static && bower install --allow-root

RUN go install ria

VOLUME ["/gopath/src/ria"]

EXPOSE 8080
EXPOSE 80
EXPOSE 8081

WORKDIR /gopath/src/ria

CMD ["ria"]

