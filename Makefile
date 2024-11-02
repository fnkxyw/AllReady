
# up all migrations
goose-up:
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up-by-one
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up-by-one
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up-by-one
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up-by-one
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up-by-one
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up-by-one
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up-by-one
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up-by-one

# down all migrations
goose-down:
	goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" reset

