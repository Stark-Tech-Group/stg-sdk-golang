FROM starktechgroup/stark-index-service:latest

USER root

RUN apk update && apk add curl
