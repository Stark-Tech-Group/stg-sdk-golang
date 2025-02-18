FROM starktechgroup/notification-service:latest

USER root
RUN apk update && apk add curl