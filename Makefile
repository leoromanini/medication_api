main_package_path = ./cmd/web

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## test: run all tests (no caching)
.PHONY: test
test:
	go test -v ./... -count=1

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## run: run all the services on Docker
.PHONY: run
run:
	docker compose up -d

## run-build: run all the services on Docker forcing build to run
.PHONY: run-build
run-build:
	docker compose down
	docker compose up --build -d

## stop: stop all the running Docker services
.PHONY: stop
stop:
	docker compose down

## teardown: stop all the running Docker services deleting volumes
.PHONY: teardown
teardown:
	docker compose down -v

## run: run the  application locally (outside Docker)
.PHONY: run-locally
run-locally: tidy
	docker compose up mysql -d

	until [ "$$(docker inspect --format='{{.State.Health.Status}}' $$(docker-compose ps -q mysql))" = "healthy" ]; do \
		echo "Waiting for MySQL to be healthy..."; \
		sleep 1; \
	done
	sleep 5;

	go run ${main_package_path} -dsn "web:password@tcp(localhost:3306)/healthcare?parseTime=true"
