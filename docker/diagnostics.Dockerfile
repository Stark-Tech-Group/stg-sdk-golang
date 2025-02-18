FROM starktechgroup/stark-diagnostics:latest

USER root

RUN apk update && apk add curl
