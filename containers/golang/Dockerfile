FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

RUN \
  mkdir -p /goroot && \
  curl https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

RUN \
    mkdir /gopath

RUN apt-get install -y --no-install-recommends mercurial

ENV GOROOT /goroot
ENV GOPATH /gopath
ENV GOBIN $GOROOT/bin
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

# Define working directory.
WORKDIR /gopath
