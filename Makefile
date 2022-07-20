test:
	go get -t -v ./...
	go test -v ./...

test_ci:
	docker-compose build --no-cache app
	docker-compose run --rm app make test

