#!/bin/bash

# MySQLサーバーをバックグラウンドで起動
mysqld_safe &

# MySQLが起動するまで待機
until mysql -h127.0.0.1 -uroot -e "SELECT 1"; do
  echo "Waiting for MySQL to start..."
  sleep 3
done

# MySQLが起動したらアプリケーションを実行
echo "MySQL is up, starting the app..."
./build/main