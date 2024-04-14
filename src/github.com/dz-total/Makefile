run:
	go run .

test:
	docker-compose -f docker-compose.yml up -d
	go test -v ./dbs
	go test -v ./handlers
	docker-compose -f docker-compose.yml down


test-db-up:
	docker-compose -f docker-compose.yml up -d

test-db-down:
	docker-compose -f docker-compose.yml down

clean:
	go clean