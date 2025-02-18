FROM starktechgroup/stark-pjm-data-pump:latest

USER root
RUN apk update && apk add curl
