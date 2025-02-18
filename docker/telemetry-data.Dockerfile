FROM starktechgroup/stark-telemetry-data-service:latest

USER root

RUN apk update && apk add curl
