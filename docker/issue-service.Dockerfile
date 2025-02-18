FROM starktechgroup/stark-issue-service:latest

USER root

RUN apk update && apk add curl
