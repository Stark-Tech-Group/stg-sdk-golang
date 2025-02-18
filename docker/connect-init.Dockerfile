FROM starktechgroup/kafka-connect-init

USER root

RUN apk update && apk add curl
