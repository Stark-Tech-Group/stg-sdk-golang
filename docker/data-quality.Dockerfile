FROM starktechgroup/stark-data-quality-service:latest

USER root
RUN apk update && apk add curl
