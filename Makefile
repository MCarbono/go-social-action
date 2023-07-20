run:
	go run main.go

tests:
	go test ./... -v

test_unit:
	go test ./domain/entity -v

test_integration:
	go test ./test/integration -v

docker_down:
	docker compose down --remove-orphans

docker_up:
	docker compose up -d