FROM starktechgroup/stark-yahoo-data-pump:latest

USER root

RUN apk update && apk add curl
