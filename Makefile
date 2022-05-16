.DEFAULT_GOAL := help

help:
	@cat $(MAKEFILE_LIST) | \
	    perl -ne 'print if /^\w+.*##/;' | \
	    perl -pe 's/(.*):.*##\s*/sprintf("%-20s",$$1)/eg;'

fmt: FORCE ## Format Code
	goimports -w ./

lint: FORCE ## Lint Code
	go vet ./...

test: FORCE ## Run Test
	go test -v -cover ./...

FORCE: