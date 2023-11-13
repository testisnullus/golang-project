DIR=${CURDIR}

migrations-up:
	docker run --rm -v $(DIR)/pkg/users/repository/migrations:/migrations \
	--network golang-project_default migrate/migrate -path=/migrations/ \
		-database postgres://dev:12345@users-postgres:5432/users?sslmode=disable up


dev:
	docker-compose up -d --build users-postgres
	docker-compose up -d --build weather-postgres
	docker-compose up -d --build weather
	docker-compose up -d --build users
	make migrations-up
	make migrations-up-weather
	docker-compose up --build traefik

migrations-up-weather:
	docker run --rm -v $(DIR)/pkg/weather/repository/migrations:/migrations \
	--network golang-project_default migrate/migrate -path=/migrations/ \
		-database postgres://dev:12345@weather-postgres:5432/weather?sslmode=disable up
