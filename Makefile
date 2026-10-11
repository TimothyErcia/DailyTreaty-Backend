main_package_path = .
binary_name = RNTreatyBackend

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
