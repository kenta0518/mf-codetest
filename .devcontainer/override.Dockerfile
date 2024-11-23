FROM mcr.microsoft.com/vscode/devcontainers/go:1.20-bullseye

WORKDIR /app

ENV TZ=Asia/Tokyo

# sudoをインストール（必要であれば）
USER root
RUN apt-get update && apt-get install -y sudo

# MySQLのGPG公開鍵を追加
RUN wget https://dev.mysql.com/doc/refman/8.0/en/checking-gpg-signature.html | tee /etc/apt/trusted.gpg.d/mysql.asc
RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys B7B3B788A8D3785C

# MariaDBのリポジトリを追加
RUN wget https://dev.mysql.com/get/mysql-apt-config_0.8.17-1_all.deb \
    && dpkg -i mysql-apt-config_0.8.17-1_all.deb \
    && apt-get update

# MariaDBのインストール
RUN apt-get install -y mariadb-server mariadb-client \
    && rm -rf /var/lib/apt/lists/*

# Goモジュールの依存関係のインストール
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

# アプリケーションのコピー
COPY ../ ./

# MariaDBの初期化
RUN mysql_install_db --user=mysql --basedir=/usr --datadir=/var/lib/mysql

# アプリケーションのファイルのオーナーを変更
RUN chown -R vscode:vscode /go

# MariaDBサーバーをバックグラウンドで起動し、アプリケーションを実行
CMD mysqld_safe & ./build/main