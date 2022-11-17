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

make test:
	@go test ./...