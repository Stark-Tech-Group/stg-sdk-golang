FROM starktechgroup/stark-tag-api:latest

USER root

RUN apk update && apk add curl

