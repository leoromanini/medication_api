# Go Medications REST API
This project focuses on the design and implementation of a RESTful API using Go to manage medication records.

It is a example for making simple RESTful API with Go using [go-chi/chi](https:-github.com/go-chi/chi), a lightweight, idiomatic and composable router for building Go HTTP services.

## Requirements
- [git](https:-git-scm.com/downloads)
- [docker engine](https:-docs.docker.com/engine/install/) >= v27.0.3
- [docker compose](https:-docs.docker.com/compose/install/) >= v2.28.1
- [make](https:-www.gnu.org/software/make/) - GNU Make 4.3
- [Golang](https:-go.dev/doc/install) >= 1.23.4 - For running the app outside Docker.

**Note:** The versions listed above were previously tested with this codebase and can be used as a reference, but **is not required for you to have them exactly**. Older versions might lead to unexpected issues though.

## Running all the services on Docker:
```bash
git clone git@github.com:leoromanini/medication_api.git
cd medication_api
make run
```
### Services list
- web (Golang REST API)
- mysql
- prometheus

## API Endpoint : http://localhost:4000

#### /medications
* `GET`     : Get all medications
* `POST`    : Create a new medication

#### /medications/:id
* `GET`     : Get a medications
* `PATCH`   : Update a medications
* `DELETE`  : Delete a medications

#### /health
* `GET` : Web app healthcheck

#### /metrics
* `GET` : Raw metrics outputs used by Prometheus

**Note:** You can check some example requests in raw [cURL](./_examples/curls.md) or [Postman Collection](./_examples/medications.postman_collection.json) in [_examples](./_examples/) folder.

## More commands

### Running all the tests (unity and integration):
```bash
make test
```
**Note:** Run `make run` before `make test` if you want to test integrations.

### Running tests + quality control checks:
```bash
make audit
```

### Running tests + coverage:
```bash
make test/cover
```

### Rebuild Dockerfile and start all the services again:
```bash
make run-build
```

### Stopping all the services on Docker:
```bash
make stop
```

### Clearing your environment:
```bash
make teardown   # Same as "make stop" + deleting docker volumes
```

### Running API service locally (outside Docker):
```bash
make run-locally   # Will start web app locally + docker mysql
```
**Note:** I recommend to use this method for daily basis delevoplemt instead of `make run-build` and `make run` for each execution.

## Structure
```
├── cmd/web                     - application-specific code path (related to business rules)
│   ├── handlers.go                 - handlers definitons
│   ├── handlers_test.go            - handlers unity tests
│   ├── middlewares.go              - middlewares definitons
│   ├── middlewares_test.go         - middlewares unity tests
│   ├── main.go                     - app entrypoint
│   ├── routes.go                   - routes definition
│   ├── validator.go                - business rules validations
│   ├── testutils_test.go           - unity tests utils
│   └── helpers.go                  - status codes helpers
├── internal/models             - non-application-specific code path (could potentially be reused)
│   ├── medications.go              - model definition
│   ├── medications_test.go         - model integration test
│   ├── errors.go                   - model specific errors
│   ├── testutils_test.go           - integration tests utils
│   ├── mocks                   - mocks path
│   │   └── medications.go          - model functions mocks
│   ├── testdata                - testdata path
│   │   ├── setup.sql               - SQL script used by integration test startup
│   │   └── teardown.sql            - SQL script used by integration test teardown
├── docker-compose.yaml         - docker compose service definitions
├── Dockerfile                  - go web app Dokerfile
├── init.sql                    - script used to create table, users and ingest sample initial data on DB
├── prometheus.yaml             - prometheus config file
└── _examples                   - examples path
    ├── curls.md                - cURL basic examples of API's usage
    └── *.postman_collection    - postman colletion with some basic examples of API's usage
```

## Next steps / Possible improvements
- Pagination.
- Support Authentication.
- Add Swagger.
- Add E2E tests.
- Support for HTTPS/TLS.
- Add a caching layer.