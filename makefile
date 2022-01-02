run:
	go run cmd/main.go

test:
	go test ./...

test_coverage:
	go test -cover ./...