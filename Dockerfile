FROM golang:1.17.3-alpine3.14 as builder
ENV PROJECT informado
RUN mkdir $PROJECT && \
    adduser -D -g '' $PROJECT
COPY cmd ./$PROJECT/cmd/
COPY internal ./$PROJECT/internal/
COPY go.mod go.sum ./$PROJECT/
WORKDIR $PROJECT/cmd/$PROJECT
RUN apk add git && \
    CGO_ENABLED=0 go build && \
    cp $PROJECT /$PROJECT

FROM alpine:3.14.3
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /informado /usr/local/bin/informado
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
USER informado
ENTRYPOINT ["informado"]
