FROM clojure:openjdk-17-lein-2.9.6-slim-buster as deps

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY project.clj /usr/src/app/

RUN lein deps

COPY . /usr/src/app

RUN mv "$(lein uberjar | sed -n 's/^Created \(.*standalone\.jar\)/\1/p')" app-standalone.jar

FROM openjdk:11.0.11-jre-slim-buster

# ENV IBKR_GATEWAY_TARGET="http://example.com"
ENV HTTP_CLIENT_TIMEOUT_MS="5000"
ENV HTTP_SOCKET_TIMEOUT_MS="5000"

COPY --from=deps /usr/src/app/app-standalone.jar /usr/bin/app-standalone.jar
COPY ./tools/repeat.sh /tmp/repeat.sh

CMD ["java", "-jar", "/usr/bin/app-standalone.jar"]