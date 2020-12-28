#!/usr/bin/env bash

SCRIPTPATH="$(cd "$(dirname "$0")"; pwd -P)"

echo $SCRIPTPATH

echo "[Restore] MySQL"
rm -rf $SCRIPTPATH/*.sql
tar zxf $SCRIPTPATH/listenfield.tar.gz -C $SCRIPTPATH/
docker run --rm -it \
--network=listenfield_local \
--volume=$SCRIPTPATH/listenfield_2020-12-29.sql:/tmp/mysql_seed.sql \
mysql:5.7 \
bash -c 'mysql -h mysql -uroot -proot listenfield < /tmp/mysql_seed.sql'
echo ""