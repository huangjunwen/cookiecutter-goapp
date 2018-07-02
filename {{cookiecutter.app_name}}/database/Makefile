
# Generate models code.
models: mysql_server_started
	@sqlw-mysql -dsn "root:123456@tcp(localhost:{{cookiecutter.devdb_port}})/dev?parseTime=true" \
		-out ../models -stmt stmts

# Launch a dev MySQL client.
mysql_client: mysql_server_started
	@docker exec -it "devdb_{{cookiecutter.app_name}}" mysql -uroot -p123456 dev


# Checks whether the dev MySQL server is launched.
mysql_server_started:
	@docker container inspect "devdb_{{cookiecutter.app_name}}" > /dev/null 2>&1 || \
		{ echo >&2 "Dev db server not started. Run 'make mysql_server' first"; exit 1; }


# Launch a dev MySQL server and initialize with current migrations.
mysql_server:
	@docker container inspect "devdb_{{cookiecutter.app_name}}" > /dev/null 2>&1 || ( \
		mkdir -p .initdb.d && \
		python -c "from schemas import create_all; print create_all()" > .initdb.d/db.sql && \
		docker run --rm --name "devdb_{{cookiecutter.app_name}}" -p 127.0.0.1:{{cookiecutter.devdb_port}}:3306 \
    -v $(shell pwd)/.initdb.d:/docker-entrypoint-initdb.d \
    -e MYSQL_ROOT_PASSWORD=123456 \
    -e MYSQL_DATABASE=dev \
    mysql:5.7.21 )


