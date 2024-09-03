FROM openjdk:8u212-jre-alpine3.9

ARG USERNAME="ibkr-cp"
ARG USERID=777

RUN addgroup -g ${USERID} -S ${USERNAME} && \
    adduser -g ${USERID} -S ${USERNAME} -G ${USERNAME} && \
    mkdir /home/${USERNAME}/portal

COPY --chown=${USERNAME} . /home/${USERNAME}/portal/

RUN chown -R ${USERNAME}:${USERNAME} /home/${USERNAME}/portal && \
    chmod -R 700 /home/${USERNAME}/portal

USER ${USERNAME}

WORKDIR /home/${USERNAME}/portal

EXPOSE 35001

ENTRYPOINT [ "/bin/sh", "bin/run.sh", "root/conf.yaml" ]