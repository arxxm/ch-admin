FROM git-registry.pingocean.com/docker-images/alpine:3.8

COPY clickhouse-admin-sms-consumer .
ADD configs/. configs/.

RUN apk --no-cache add openssl curl tzdata && \
    chmod a+x clickhouse-admin-sms-consumer

ENTRYPOINT [ "/clickhouse-admin-sms-consumer" ]