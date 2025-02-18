FROM starktechgroup/stark-babbage:latest

USER root

RUN apk update && apk add curl
