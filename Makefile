build-all:
	cd goapi && GOOS=linux GOARCH=amd64 make build

run-all: build-all
	docker-compose up --force-recreate --build

precommit:
	cd goapi && make precommit


PHONY: proto-generate
proto-generate:
	cd goapi && make proto-generate

.PHONY: proto-fast-generate
proto-fast-generate:
	cd goapi && make proto-fast-generate