migrate-up:
	migrate -path internal/migration/ -database "postgresql://postgres:root@localhost:5432/postgres?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/migration/ -database "postgresql://postgres:root@localhost:5432/postgres?sslmode=disable" -verbose down