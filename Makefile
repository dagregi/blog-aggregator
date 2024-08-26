BINARY=aggregator
BIN=bin

SCHEMA:=sql/schema
DB:=$(shell grep DATABASE_URL .env | sed 's/DATABASE_URL=//; s/?.*//')

all: clean test build

clean:
	@echo cleaning...
	@rm -f ${BIN}/${BINARY}

build:
	@go build -o ${BIN}/${BINARY}

run: build
	./${BIN}/${BINARY}

test:
	@go test ./... -v

migrate/up:
	@goose -dir ${SCHEMA} postgres ${DB} up

migrate/down:
	@goose -dir ${SCHEMA} postgres ${DB} down
