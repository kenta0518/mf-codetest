#!/bin/bash

# MariaDBをバックグラウンドで起動
service mariadb start

# MariaDBの初期設定
mysql -u root <<EOF
ALTER USER 'root'@'localhost' IDENTIFIED VIA mysql_native_password;
FLUSH PRIVILEGES;
CREATE DATABASE IF NOT EXISTS codetest;
EOF

# サービスをフォアグラウンドで維持（必要であれば）
tail -f /dev/null