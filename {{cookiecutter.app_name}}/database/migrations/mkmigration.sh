#!/bin/bash

# 该脚本自动根据当前 schemas 内定义的数据表结构与 migrations 中的差异计算出新的数据库迁移版本

if [ -z "$1" ]; then
  echo "Usage: mkmigration.sh \"migration description\""
  exit 1
fi

if [ "$(docker ps -q -f name=migrationdb_{{cookiecutter.app_name}})" ]; then
  echo ">>>>> Migration is running. Aborting."
  exit 1
fi

echo ">>>>> Starting migration db ..."
docker run -d --rm --name "migrationdb_{{cookiecutter.app_name}}" -p 127.0.0.1:26033:3306 \
  -e MYSQL_ROOT_PASSWORD=123456 \
  -e MYSQL_DATABASE=migration \
  mysql:5.7.21

function fin {
  echo ">>>>> Killing migration db ..."
  docker kill "migrationdb_{{cookiecutter.app_name}}"
  echo ">>>>> Killed."
}
trap fin EXIT

# https://stackoverflow.com/a/48703384/157235
while ! docker exec "migrationdb_{{cookiecutter.app_name}}" mysql -uroot -p123456 -e "status" > /dev/null 2>&1 ; do
  echo ">>>>> Waiting 2 seconds for migration db up ..."
  sleep 2
done

echo ">>>>> Loading current db schema ..."
alembic -c ../migrations/alembic.ini upgrade head

echo ">>>>> Generating new migration ..."
alembic -c ../migrations/alembic.ini revision --autogenerate -m "$1"
