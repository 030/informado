FROM golang:1.17.3-alpine3.14 as builder
ENV USERNAME=informado
ENV BASE=/opt/${USERNAME}
COPY . ./${USERNAME}/
WORKDIR /go/${USERNAME}/cmd/${USERNAME}/
RUN apk add curl git && \
    CGO_ENABLED=0 go build && \
    curl -sL https://gist.githubusercontent.com/030/54fc7ae735a163c09dcf6f3699d87e81/raw/b82514f50525e0ebf843e0dbf9bef1a382ccd40f/openshift-docker-user-entrypoint.sh > entrypoint.sh && \
    curl -sL https://gist.githubusercontent.com/030/34a2bf3f7f1cd427dc36c86dcb1e8cf7/raw/e1be7ef3c2c1a8441e406a669a0f6b6d97dcc984/openshift-docker-user.sh > user.sh && \
    chmod +x user.sh && \
    ./user.sh

FROM alpine:3.14.3
ENV BIN=/usr/local/bin/
ENV USERNAME=informado
ENV BASE=/opt/${USERNAME}
ENV BASE_BIN=${BASE}/bin
ENV PATH=${BASE_BIN}:${PATH}
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /opt/ /opt/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER informado
ENTRYPOINT ["entrypoint.sh"]
CMD ["informado"]
