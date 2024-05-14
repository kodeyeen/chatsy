postgres:
	docker compose exec -it postgres psql -U postgres -d postgres

createdb:
	docker compose exec -it postgres createdb -U postgres -O postgres chatsy

dropdb:
	docker compose exec -it postgres dropdb -U postgres chatsy

migrateup:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/chatsy?sslmode=disable' up

migratedown:
	migrate -path ./migrations -database 'postgres://postgres:postgres@localhost:5432/chatsy?sslmode=disable' down

.PHONY: migrateup migratedown

# migrate create -seq -ext=".sql" -dir="./migrations" create_users_table
