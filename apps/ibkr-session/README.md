### Overview

This is a service that has the following responsibilities:
- Attempts to keep the IBKR session open and authentication refreshed
- Reports authentication failures in logging

For more information see ``` refresh/doc/intro.md ```

### Clojure HowTo

New Project in Leinengen
```
PRJ_NAME="refresh"
docker run --entrypoint "lein" --rm --volume $(PWD):/tmp --workdir /tmp clojure:openjdk-17-lein-2.9.6-slim-buster new app ${PRJ_NAME}
```

### Usage

* Note the following required environment variables
    * IBKR_GATEWAY_TARGET
        * The URL to the ibkr-gateway
    * HTTP_CLIENT_TIMEOUT_MS="5000"
        * Timeout in MS for http connections from this app
    * HTTP_SOCKET_TIMEOUT_MS="5000"
        * Timeout in MS for TCP sockets from this app

* Build the docker image

```
/bin/sh docker-build.sh
```

* Execute the docker image 

```
docker run --env IBKR_GATEWAY_TARGET="http://example.com" --env HTTP_CLIENT_TIMEOUT_MS=5000 --env HTTP_SOCKET_TIMEOUT_MS=5000 localhost:5000/ibkr-session:latest
```
