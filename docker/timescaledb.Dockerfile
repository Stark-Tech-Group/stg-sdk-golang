ARG PG_MAJOR=12
FROM debezium/postgres:$PG_MAJOR AS debezium

ENV POSTGRES_HOST_AUTH_METHOD=trust

# FROM postgres:10.5-alpine AS wal2jsonbuilder
# RUN apk add --no-cache build-base git
# RUN git clone https://github.com/eulerto/wal2json.git
# WORKDIR wal2json
# RUN USE_PGXS=1 make && make install

FROM timescale/timescaledb:latest-pg$PG_MAJOR
# install debezium dependencies
COPY --from=debezium /usr/lib/postgresql/$PG_MAJOR/lib/decoderbufs.so /usr/local/lib/postgresql/
COPY --from=debezium /usr/share/postgresql/$PG_MAJOR/extension/decoderbufs.control /usr/local/share/postgresql/extension/
# COPY --from=wal2jsonbuilder /wal2json/wal2json.so /usr/local/lib/postgresql/

# build wal2json for debezium
ENV PLUGIN_VERSION=v1.5.0.Beta1
ENV WAL2JSON_COMMIT_ID=wal2json_2_3
RUN apk add --no-cache protobuf-c-dev
RUN apk add --no-cache --virtual .debezium-build-deps gcc clang llvm git make musl-dev pkgconf \
    && git clone https://github.com/debezium/postgres-decoderbufs -b $PLUGIN_VERSION --single-branch \
    && (cd /postgres-decoderbufs && make && make install) \
    && rm -rf postgres-decoderbufs \
    && git clone https://github.com/eulerto/wal2json -b master --single-branch \
    && (cd /wal2json && git checkout $WAL2JSON_COMMIT_ID && make && make install) \
    && rm -rf wal2json \
    && apk del .debezium-build-deps


