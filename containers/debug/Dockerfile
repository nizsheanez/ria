FROM nizsheanez/ubuntu
MAINTAINER Alex Sharov <www.pismeco@gmail.com>

# ZSH
RUN \
        apt-get install -y --no-install-recommends nano zsh && \
        wget https://raw.githubusercontent.com/robbyrussell/oh-my-zsh/master/tools/install.sh -O - | zsh && \
        chsh -s `which zsh`

ADD ./zsh /opt/configs/zsh

RUN \
        cp /opt/configs/zsh/.zshrc /root/.zshrc && \
        rm -rf /opt/configs/

EXPOSE 22

CMD ["zsh"]