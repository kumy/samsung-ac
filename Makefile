
COVER_OUT ?= /tmp/cover.out
COVER_HTML ?= /tmp/cover.html

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
.DEFAULT_GOAL := help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)


build: ## build the code
	@go build

run: build ## build the code
	./det

mock: ## run mockery to regenerate mocks
	mockery

test: ## run Unit Tests
	grc go test ./... -cover -coverprofile $(COVER_OUT)

#	@bash -c 'while read p || [ -n "$$p" ]; do \
#		sed -i "/$${p//\//\\\/}/d" $(COVER_OUT); \
#	done < ./exclude-from-code-coverage.txt'

	go tool cover -html=$(COVER_OUT) -o $(COVER_HTML)

	@echo
	@echo "Coverage html: file://$(COVER_HTML)"
	@echo
