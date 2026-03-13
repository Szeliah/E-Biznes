FROM ubuntu:24.04

RUN apt update -y && apt upgrade -y && apt install -y curl unzip

WORKDIR /usr/local

RUN apt install -y openjdk-8-jdk

RUN curl -L https://github.com/JetBrains/kotlin/releases/download/v2.3.10/kotlin-compiler-2.3.10.zip -o kotlin.zip \
&& unzip kotlin.zip && rm kotlin.zip

ENV PATH="/usr/local/kotlinc/bin:${PATH}"

WORKDIR /

CMD ["/bin/bash"]