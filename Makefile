.PHONY: run

run:
	go run main.go

migrate-up:
	migrate -database postgres://gofiber:gofiber@localhost:5432/gofiber?sslmode=disable -path ./migrations -verbose up

migrate-down:
	migrate -database postgres://gofiber:gofiber@localhost:5432/gofiber?sslmode=disable -path ./migrations -verbose down
