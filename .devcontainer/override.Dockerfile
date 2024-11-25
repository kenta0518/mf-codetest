FROM mcr.microsoft.com/vscode/devcontainers/go:1.20-bullseye

WORKDIR /app

ENV TZ=Asia/Tokyo

# 必要なパッケージをインストール
USER root
RUN apt-get update && apt-get install -y \
    sudo \
    mariadb-server mariadb-client \
    && rm -rf /var/lib/apt/lists/*

# MySQLデータディレクトリを初期化
RUN mysql_install_db --user=mysql --basedir=/usr --datadir=/var/lib/mysql

# MariaDB初期設定用スクリプトを作成
COPY init-mariadb.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/init-mariadb.sh

# Goモジュールの依存関係をインストール
COPY ../go.mod ../go.sum ./
RUN go mod download

# 必要なGoツールをインストール
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.0
RUN go install honnef.co/go/tools/cmd/staticcheck@v0.4.0
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/golang/mock/mockgen@v1.7.0-rc.1
RUN go install github.com/swaggo/swag/cmd/swag@latest

# アプリケーションコードをコピー
COPY ../ ./

# ENTRYPOINTスクリプトを作成
COPY entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["entrypoint.sh"]