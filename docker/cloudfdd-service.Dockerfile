FROM starktechgroup/cloudfdd-service:latest 

USER root

RUN apk update && apk add curl
