FROM starktechgroup/stark-nyiso-data-pump:latest

USER root

RUN apk update && apk add curl
