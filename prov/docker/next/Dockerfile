FROM node:alpine

ARG project_dir=/opt/app
WORKDIR $project_dir

RUN set -eux \
  # Linux Package Install
  && apk --update add \
    # for Develop
    git \
    # for Test
    curl \
    # node-gyp ビルド用のインストール for M1 Mac
    python3 \
    make \
    g++ \
    vips-dev \
    # Timezone
    tzdata \
  # Change Timezone
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  # clean up
  && apk del tzdata \
  && rm -rf /var/cache/apk/* \
  # typesync
  ## 必要なtypescriptの型定義をdevDependencies側に入れてくれる
  && yarn global add typesync
