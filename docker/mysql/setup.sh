#!/bin/bash

DATADIR="/data/mysql"

echo 'Running mysql_install_db ...'
mysql_install_db --datadir=$DATADIR
echo 'Finished mysql_install_db'

tempSqlFile='/tmp/mysql-first-time.sql'

cat > "$tempSqlFile" <<-EOSQL
DELETE FROM mysql.user ;
CREATE USER 'root'@'%' IDENTIFIED BY 'root' ;
GRANT ALL ON *.* TO 'root'@'%' WITH GRANT OPTION ;
CREATE USER 'myadmin'@'%' IDENTIFIED BY 'root' ;
GRANT ALL ON *.* TO 'myadmin'@'%' WITH GRANT OPTION ;
DROP DATABASE IF EXISTS test ;
EOSQL

echo 'FLUSH PRIVILEGES ;' >> "$tempSqlFile"

chown -R mysql:mysql "$DATADIR"

mysqld --datadir=$DATADIR --user=mysql --pid-file=/data/mysql/mysql.pid --sql-mode="" --init-file=$tempSqlFile &

while true; do
    mysql --password=root -e 'show databases;'
    if [ $? -eq 0 ]; then
        break
    else
        sleep 2
    fi
done
set -e


echo 'Building mapmyfitness'
mysql --password=root -e 'CREATE DATABASE mapmyfitness;'
# This is dumped from mysql06
echo 'Building mapmyfitness structure'
mysql --password=root mapmyfitness < /sql/mapmyfitness.structure.sql
# This is managed locally
echo 'Adding mapmyfitness data'
mysql --password=root mapmyfitness < /sql/mapmyfitnesss.dev.content.sql


kill -TERM $(cat /data/mysql/mysql.pid)
wait