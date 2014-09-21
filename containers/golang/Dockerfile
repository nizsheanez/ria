FROM google/golang
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

RUN \
    apt-get update && apt-get install -y --no-install-recommends mercurial && \
    git clone https://github.com/pote/gvp.git && \
    cd gvp && \
    git checkout v0.1.0 && \
    ./configure && make install && \
    rm -rf * && \
    git clone https://github.com/pote/gpm.git && \
    cd gpm && \
    git checkout v1.2.3 && \
    ./configure && make install && \
    rm -rf *

ENV GOBIN $GOROOT/bin

