start-db:
	@docker run -d \
		--name spots \
		-p 5432:5432 \
		-e POSTGRES_PASSWORD=complexpassword \
		postgis/postgis
	@docker cp ./docker-entrypoint-initdb.d/spots.sql spots:/docker-entrypoint-initdb.d/spots.sql

test-db:
	@docker run -d \
		--name spots_test \
		-p 5433:5432 \
		-e POSTGRES_PASSWORD=VerySecurePassword \
		postgis/postgis
		@docker cp ./docker-entrypoint-initdb.d/spots.sql spots_test:/docker-entrypoint-initdb.d/spots.sql


cleanup:
	@docker stop spots
	@docker rm spots
	@docker stop spots_test
	@docker rm spots_test


test:
	@go test ./...