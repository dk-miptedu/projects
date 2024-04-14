run:
	go run .

test:
	docker-compose -f docker-compose.yml up -d
	go test -v ./db
	go test -v ./handlers
	go test -v ./cmd
	go test -v ./models
	docker-compose -f docker-compose.yml down


test-db-up:
	docker-compose -f docker-compose.yml up -d

test-db-down:
	docker-compose -f docker-compose.yml down

clean:
	go clean