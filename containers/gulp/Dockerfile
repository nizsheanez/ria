FROM dockerfile/nodejs

RUN apt-get update
RUN apt-get install -y ruby-dev build-essential

RUN npm install -g gulp

RUN mkdir /gulp

WORKDIR /gopath/src/ria/static

RUN

CMD ["gulp", "watch"]
