FROM starktechgroup/stark-search-service:latest

USER root

RUN apk update && apk add curl
