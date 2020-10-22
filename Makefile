APP=go-store
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor" | grep -v "docs")
IMPORTS_PACKAGES=$(shell go list ./... | grep -v "vendor" | grep -v "docs" | awk -F"\/" '{print $$4}')
APP_EXECUTABLE="./out/$(APP)"

setup:
	go get -u golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports

lint:
	@echo "Running lint..."
	@for p in $(ALL_PACKAGES); do \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

imports:
	@goimports -w -local http://github.com/yaminmhd/go-store-api-docker $(IMPORTS_PACKAGES)

vet:
	@echo "Running vet..."
	@go vet ./...

compile:
	@echo "Building executable..."
	@mkdir -p out/
	@go build -o $(APP_EXECUTABLE)

build: imports lint vet compile

test:
	@go test $(ALL_PACKAGES)

run: compile
	@echo "Running server..."
	@$(APP_EXECUTABLE) start
