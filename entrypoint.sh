#!/bin/bash

# MariaDBの起動
mysqld_safe &
sleep 5 # 起動を待つ

# 初期化スクリプトの実行
init-mariadb.sh

# アプリケーションの起動
exec ./build/main