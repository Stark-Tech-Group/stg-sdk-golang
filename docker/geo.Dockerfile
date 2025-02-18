FROM starktechgroup/stg-geo:latest

USER root

RUN apk update && apk add curl
