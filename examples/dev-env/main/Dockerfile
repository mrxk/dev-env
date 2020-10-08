FROM golang

ENV GOOS=linux \
    GOARCH=arm \
    GOARM=5 \
    SHELL=/bin/bash

ARG USER_ID=1000
ARG GROUP_ID=1000

RUN apt-get update && \
    apt-get install -y sudo tmux vim && \
    apt-get clean

RUN groupadd -g ${GROUP_ID} user1 && \
    useradd -l -u ${USER_ID} -g user1 user1 && \
    usermod -aG sudo user1 && \
    install -d -m 0755 -o user1 -g user1 /home/user1 && \
    echo "user1 ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/010_user1-nopasswd

USER user1

COPY tmux.conf /home/user1/.tmux.conf
COPY bashrc /home/user1/.bashrc

RUN echo source ~/.bashrc > ~/.bash_profile

RUN git config --global user.email "user1@user1" && \
    git config --global user.name "user1"

ENTRYPOINT ["tmux"]
