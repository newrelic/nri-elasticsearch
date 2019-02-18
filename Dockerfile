FROM golang:1.9 as builder
RUN go get -d github.com/newrelic/nri-elasticsearch/... && \
    cd /go/src/github.com/newrelic/nri-elasticsearch && \
    make && \
    strip ./bin/nr-elasticsearch

FROM newrelic/infrastructure:latest
COPY --from=builder /go/src/github.com/newrelic/nri-elasticsearch/bin/nr-elasticsearch /var/db/newrelic-infra/newrelic-integrations/bin/nr-elasticsearch
COPY --from=builder /go/src/github.com/newrelic/nri-elasticsearch/elasticsearch-definition.yml /var/db/newrelic-infra/newrelic-integrations/definition.yml
