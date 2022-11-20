start:
	@docker run -d \
		--name spots \
		-p 5432:5432 \
		-e POSTGRES_PASSWORD=complexpassword \
		postgis/postgis
#		-v $(PWD)/docker-entrypoint-initdb.d/:/docker-entrypoint-initdb.d/ \
#		postgis/postgis

make cleanup:
	@docker stop spots
	@docker rm spots

make test-db:
	@docker run -d \
		--name spots_test \
		-p 5433:5432 \
		-e POSTGRES_PASSWORD=VerySecurePassword \
		postgis/postgis

make test:
	@go test ./...