dirname=$(shell basename $(CURDIR))
get-dir-name:
	echo $(dirname)

prod:
	go run main.go

start: start-containers
	go run main.go

clean-containers:
	docker-compose kill
	docker-compose rm -f

start-containers: clean-containers
	docker-compose up -d
	make wait-for-db

wait-for-db:
	docker run --rm --network '$(dirname)_default' busybox /bin/sh -c "until nc -z db 3306; do sleep 3; echo 'Waiting for DB to be available...'; done"

test:
	go test ./...

test.integration: start-containers
	go test ./integration -count=1
