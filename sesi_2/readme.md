# How to migrate database

- up: `$ migrate -path database/migration/ -database "postgresql://postgres:afif123@localhost:5432/postgres?sslmode=disable" -verbose up`
- down: `$ migrate -path database/migration/ -database "postgresql://postgres:afif123@localhost:5432/postgres?sslmode=disable" -verbose up`