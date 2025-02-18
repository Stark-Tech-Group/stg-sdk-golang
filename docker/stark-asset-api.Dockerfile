FROM starktechgroup/stark-asset-api:latest
USER root
RUN apk update && apk add curl
