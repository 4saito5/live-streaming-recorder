FROM golang:alpine
# Reference: https://gqlgen.com/getting-started/

ARG project_dir=/go/live-streaming-recorder
WORKDIR $project_dir

RUN set -eux \
  # Linux Package Install
  && apk --update add \
    # Timezone
    tzdata \
    git \
    alpine-sdk \
    sqlite \
  # Change Timezone
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  && rm -rf /var/cache/apk/*

COPY ./src/back/ ./
RUN go test
