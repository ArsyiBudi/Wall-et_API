postgres:
	docker run --name postgres-12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres:12-alpine
createdb:
	 docker exec -it postgres-12 createdb --username=root --owner=root walletBank

dropdb:
	  docker exec -it postgres-12 dropdb walletBank

migrateup:
		migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/walletBank?sslmode=disable" -verbose up

migratedown:
		 migrate -path db/migration -database "postgresql://root:mysecretpassword@localhost:5432/walletBank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test