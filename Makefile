DSN = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"


# up all migrations
goose-up:
	goose -dir ./migrations postgres $(DSN) up-by-one
	goose -dir ./migrations postgres $(DSN) up-by-one
	goose -dir ./migrations postgres $(DSN) up-by-one
	goose -dir ./migrations postgres $(DSN) up-by-one
	goose -dir ./migrations postgres $(DSN) up-by-one
	goose -dir ./migrations postgres $(DSN) up-by-one
	goose -dir ./migrations postgres $(DSN) up-by-one
	goose -dir ./migrations postgres $(DSN) up-by-one
	goose -dir ./migrations postgres $(DSN) up-by-one

# down all migrations
goose-down:
	goose -dir ./migrations postgres $(DSN) reset

# look status of migrations
goose-status:
	goose -dir ./migrations postgres $(DSN) status

# filling db with fake data
fake-data:
	go run cmd/faker/faker.go