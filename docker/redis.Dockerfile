FROM redis:6.2.7
USER root

RUN apt-get update
RUN apt-get install -y gettext-base

COPY docker/redis.conf /etc/redis/redis.conf.template
COPY docker/start.sh /start.sh

ENTRYPOINT [ "/bin/bash", "/start.sh" ]