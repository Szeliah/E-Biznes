# Base image - in my case it is a OS Ubuntu 24.04  
FROM ubuntu:24.04

RUN apt update -y && apt install -y build-essential libssl-dev \
zlib1g-dev libffi-dev libbz2-dev libreadline-dev libsqlite3-dev curl xz-utils

WORKDIR /usr/src

RUN curl -O https://www.python.org/ftp/python/3.10.20/Python-3.10.20.tgz && tar xvf Python-3.10.20.tgz && \
cd Python-3.10.20 && ./configure --with-ensurepip=install && make altinstall && \
cd .. && rm -rf Python-3.10.20 Python-3.10.20.tgz

WORKDIR /

CMD ["/bin/bash"]