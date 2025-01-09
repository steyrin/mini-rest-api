run:
	docker-compose up
	go run cmd/main.go
	docker-compose down
