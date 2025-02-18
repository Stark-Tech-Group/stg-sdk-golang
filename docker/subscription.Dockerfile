FROM starktechgroup/subscription:latest

USER root
RUN apk update && apk add curl
