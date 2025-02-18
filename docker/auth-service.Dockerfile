FROM starktechgroup/stark-auth-service:latest

USER root
RUN apk update && apk add curl
