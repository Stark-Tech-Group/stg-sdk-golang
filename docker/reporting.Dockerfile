FROM starktechgroup/stark-reporting-service:latest
USER root
RUN apk update && apk add curl
