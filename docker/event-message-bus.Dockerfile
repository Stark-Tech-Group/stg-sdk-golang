FROM starktechgroup/stark-event-message-bus:latest

USER root

RUN apk update && apk add curl
