# 2023年3月時点の最新版goの軽量版「alpine」
FROM golang:1.20.2-alpine

# インストール可能なパッケージ一覧の更新
RUN apk update && \
    apk upgrade && \
    # パッケージのインストール（--no-cacheでキャッシュ削除）
    apk add --no-cache \
            git \
            gcc \
            musl-dev

# 作業ディレクトリの設定
WORKDIR /go/src