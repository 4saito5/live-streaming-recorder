FROM golang:alpine

ARG project_dir=/go/live-streaming-recorder
WORKDIR ${project_dir}

COPY ./prov/docker/golang/files/requirements.txt /tmp/requirements.txt

RUN set -eux \
  # Linux Package Install
  && apk --update add \
    # Timezone
    tzdata \
    git \
    alpine-sdk \
    sqlite \
    # youtube-dl
    python3 \
    py3-pip \
    # movie
    ffmpeg \
  # Change Timezone
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  && rm -rf /var/cache/apk/* \
  # Python package Install
  && pip3 install --upgrade pip \
  && pip3 install -r /tmp/requirements.txt \
  && rm -f /tmp/requirements.txt

COPY ./src/back/ ./
RUN go test
