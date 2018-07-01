
# 启动一个 MySQL dev client，devdb_server 必须已经启动
devdb_client:
	docker exec -it "devdb_{{cookiecutter.app_name}}" mysql -uroot -p123456 dev

# 启动一个 MySQL dev server，并初始化之
devdb_server:
	docker container inspect "devdb_{{cookiecutter.app_name}}" > /dev/null 2>&1 || ( \
		mkdir -p .initdb.d && \
		alembic -c migrations/alembic.ini upgrade head --sql > .initdb.d/db.sql && \
		docker run --rm --name "devdb_{{cookiecutter.app_name}}" -p 127.0.0.1:{{cookiecutter.devdb_port}}:3306 \
    -v $(shell pwd)/.initdb.d:/docker-entrypoint-initdb.d \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=dev \
    mysql:5.7.21 )


