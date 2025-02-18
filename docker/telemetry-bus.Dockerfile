FROM starktechgroup/telemetry-message-bus:latest

USER root

RUN apk update && apk add curl
