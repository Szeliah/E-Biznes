FROM ubuntu:24.04

RUN apt update -y && apt install -y curl unzip

RUN apt install -y openjdk-8-jdk openjdk-17-jdk

ENV JAVA_HOME_8=/usr/lib/jvm/java-8-openjdk-amd64 
ENV JAVA_HOME_17=/usr/lib/jvm/java-17-openjdk-amd64

WORKDIR /usr/local

RUN curl -L https://github.com/JetBrains/kotlin/releases/download/v2.3.10/kotlin-compiler-2.3.10.zip -o kotlin.zip \
&& unzip kotlin.zip && rm kotlin.zip


RUN curl -L https://services.gradle.org/distributions/gradle-9.4.0-bin.zip -o gradle.zip \
&& unzip gradle.zip && rm gradle.zip


ENV PATH="/usr/local/kotlinc/bin:${PATH}"
ENV PATH="/usr/local/gradle-9.4.0/bin:${PATH}"
ENV GRADLE_OPTS="-Dorg.gradle.java.home=${JAVA_HOME_17}"

ENV JAVA_HOME=${JAVA_HOME_8}
ENV PATH="${JAVA_HOME}/bin:${PATH}"

WORKDIR /app

COPY build.gradle /app/
RUN mkdir -p src/main/kotlin
COPY Main.kt src/main/kotlin/

RUN gradle help --no-daemon

CMD ["gradle", "run", "--no-daemon", "--quiet"]