.PHONY: help create delete info
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

create: ## create env
	sceptre launch-env example

delete: ## delete env
	sceptre delete-env example

run: ## run the example
	go run main.go

scan-cats: ## scan tables
	aws dynamodb scan --table-name Cats

scan-owners: ## scan owners table
	aws dynamodb scan --table-name Owners