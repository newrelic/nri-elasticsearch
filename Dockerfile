FROM golang:1.10 as builder
RUN go get -d github.com/newrelic/nri-elasticsearch/... && \
    cd /go/src/github.com/newrelic/nri-elasticsearch && \
    make && \
    strip ./bin/nr-elasticsearch

FROM newrelic/infrastructure:latest
ENV NRIA_IS_FORWARD_ONLY true
ENV NRIA_K8S_INTEGRATION true
COPY --from=builder /go/src/github.com/newrelic/nri-elasticsearch/bin/nr-elasticsearch /nri-sidecar/newrelic-infra/newrelic-integrations/bin/nr-elasticsearch
COPY --from=builder /go/src/github.com/newrelic/nri-elasticsearch/elasticsearch-definition.yml /nri-sidecar/newrelic-infra/newrelic-integrations/definition.yaml
USER 1000
