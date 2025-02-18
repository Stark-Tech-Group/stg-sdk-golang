FROM starktechgroup/stark-report-scheduler:latest
USER root
RUN apk update && apk add curl
