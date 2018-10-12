FROM alpine:3.8

RUN apk add --update ca-certificates \
    && rm -rf /var/cache/apk/*

ADD ./kubernetes-node-health /kubernetes-node-health

ENTRYPOINT ["/kubernetes-node-health"]
