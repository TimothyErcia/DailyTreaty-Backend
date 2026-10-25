main_package_path = .
binary_name = RNTreatyBackend

-include .env
export

## POSTGRES URL
DB_URL ?= postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
MIGRATE_DIR ?= db/migrations

# ================================================================= #
# QUALITY CONTROL
# ================================================================= #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/bin/coverage.out ./...
	go tool cover -html=/bin/coverage.out

## upgradeable: list direct dependencies that have upgrades available
.PHONY: upgradeable
upgradeable:
	go run github.com/oligot/go-mod-upgrade@latest

# ================================================================= #
# DEVELOPMENT
# ================================================================= #

## DB Migration Creation
.PHONY: create-migration
create-migration:
	@if [ -z "$(name)" ]; then echo "Error: 'name' variable is required. Example: make create-migration name=add_users_table"; exit 1; fi
	@mkdir -p db/migrations
	@TIMESTAMP=$$(date +%Y%m%d%H%M%S); \
	touch db/migrations/$${TIMESTAMP}_$(name).up.sql; \
	touch db/migrations/$${TIMESTAMP}_$(name).down.sql; \
	echo "Created db/migrations/$${TIMESTAMP}_$(name).up.sql"; \
	echo "Created db/migrations/$${TIMESTAMP}_$(name).down.sql"

.PHONY: migrate-up
# Apply all up migrations
migrate-up:
	migrate -path $(MIGRATE_DIR) -database "$(DB_URL)" -verbose up

.PHONY: migrate-down
# Roll back the last down migration (or usage: make migrate-down steps=2)
migrate-down:
	migrate -path $(MIGRATE_DIR) -database "$(DB_URL)" -verbose down $(steps)


## tidy: tidy modfiles and modernize and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fix ./...
	go fmt ./...

## build: build the application
.PHONY: build
build:
	go build -o ./bin/${binary_name} ${main_package_path}

## run: run the core(main) application
.PHONY: run
run: build
	@./bin/${binary_name}
