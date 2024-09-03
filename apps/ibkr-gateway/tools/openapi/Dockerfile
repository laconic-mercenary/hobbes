FROM swaggerapi/swagger-ui:v3.52.0

ARG USERNAME="swagger"
ARG USERID="777"

RUN addgroup -g ${USERID} -S ${USERNAME} && \
    adduser -g ${USERID} -S ${USERNAME} -G ${USERNAME}

COPY --chown=swagger:swagger ./ibkr-gateway-spec.yaml /usr/share/nginx/html/gateway.yaml
COPY --chown=swagger:swagger ./ibkr-clientportal-spec.json /usr/share/nginx/html/clientportal.json

ENV API_URL=gateway.yaml

EXPOSE 8080
