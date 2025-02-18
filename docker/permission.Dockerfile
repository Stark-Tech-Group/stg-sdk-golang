FROM starktechgroup/stark-permission-service:latest

USER root

RUN apk update && apk add curl
