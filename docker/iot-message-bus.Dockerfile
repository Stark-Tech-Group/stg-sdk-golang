FROM starktechgroup/iot-hub-message-bus:latest

USER root
RUN apk update && apk add curl
