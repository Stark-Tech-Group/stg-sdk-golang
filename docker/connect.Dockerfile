FROM starktechgroup/kafka-connect:latest
USER root
RUN confluent-hub install --no-prompt confluentinc/kafka-connect-jdbc:10.2.5
