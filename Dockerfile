FROM golang:1.10 as builder
COPY . /go/src/github.com/newrelic/nri-elasticsearch/
RUN cd /go/src/github.com/newrelic/nri-elasticsearch && \
    make && \
    strip ./bin/nri-elasticsearch

FROM newrelic/infrastructure:latest
ENV NRIA_IS_FORWARD_ONLY true
ENV NRIA_K8S_INTEGRATION true
COPY --from=builder /go/src/github.com/newrelic/nri-elasticsearch/bin/nri-elasticsearch /nri-sidecar/newrelic-infra/newrelic-integrations/bin/nri-elasticsearch
COPY --from=builder /go/src/github.com/newrelic/nri-elasticsearch/elasticsearch-definition.yml /nri-sidecar/newrelic-infra/newrelic-integrations/definition.yaml
USER 1000
