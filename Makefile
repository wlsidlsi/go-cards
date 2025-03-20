CONTAINER_NAME = postgres14
POSTGRES_USER = root
POSTGRES_DB = tweets

postgres14:
	@docker run --name $(CONTAINER_NAME) -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p 5432:5432 -d postgres:14

destroy:
	@docker stop $(CONTAINER_NAME) > /dev/null
	@docker rm $(CONTAINER_NAME) > /dev/null

createdb:
	docker exec -it $(CONTAINER_NAME) createdb -U $(POSTGRES_USER) -O $(POSTGRES_USER) $(POSTGRES_DB)

dropdb:
	docker exec -it $(CONTAINER_NAME) dropdb -U $(POSTGRES_USER) $(POSTGRES_DB)

migrateup:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@127.0.0.1:5432/$(POSTGRES_DB)?sslmode=disable" up

migratedown:
	migrate -verbose -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@127.0.0.1:5432/$(POSTGRES_DB)?sslmode=disable" down

test-live:
	USE_LIVE_DB=1 go test -v -cover ./...

test-mock:
	USE_LIVE_DB=0 go test -v -cover ./...

.PHONY:  destroy createdb dropdb postgres14 migrateup migratedown test-live test-mock