FROM golang:1.20 as builder
ARG CGO_ENABLED=0
WORKDIR /go/src/github.com/newrelic/nri-elasticsearch
COPY . .
RUN make clean compile

FROM alpine:latest
COPY --from=builder /go/src/github.com/newrelic/nri-elasticsearch/bin /
CMD ["sleep", "1h"]
