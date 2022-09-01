FROM starktechgroup/kafka-connect:latest
USER root
RUN confluent-hub install --no-prompt debezium/debezium-connector-postgresql:1.7.0
RUN confluent-hub install --no-prompt confluentinc/kafka-connect-jdbc:10.2.5
RUN confluent-hub install --no-prompt mdrogalis/voluble:0.3.1
ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
