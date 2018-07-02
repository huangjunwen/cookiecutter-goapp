
# Launch a dev MySQL client.
mysql_client: mysql_server_started
	@docker exec -it "devdb_{{cookiecutter.app_name}}" mysql -uroot -p123456 dev

# Launch a dev MySQL server and initialize with current migrations.
mysql_server:
	@docker container inspect "devdb_{{cookiecutter.app_name}}" > /dev/null 2>&1 || ( \
		mkdir -p .initdb.d && \
		alembic -c migrations/alembic.ini upgrade head --sql > .initdb.d/db.sql && \
		docker run --rm --name "devdb_{{cookiecutter.app_name}}" -p 127.0.0.1:{{cookiecutter.devdb_port}}:3306 \
    -v $(shell pwd)/.initdb.d:/docker-entrypoint-initdb.d \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=dev \
    mysql:5.7.21 )

# Checks whether the dev MySQL server is launched.
mysql_server_started:
	@docker container inspect "devdb_{{cookiecutter.app_name}}" > /dev/null 2>&1 || \
		{ echo >&2 "Dev db server not started. Run 'make db_server' first"; exit 1; }

