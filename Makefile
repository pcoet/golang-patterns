.PHONY: clean-bin build run help
.DEFAULT_GOAL := help

UHOME = ${shell echo "${HOME}"}
GOPATH = ${UHOME}/go
GOBIN = ${GOPATH}/bin

clean-bin:
	@echo "Deleting installed binary..."
	@rm -rf ${GOBIN}/testapp
build: clean-bin
	@echo "Installing to ${GOBIN}..."
	@go install github.com/pcoet/golang-patterns/cmd/testapp
run:
	@${GOBIN}/testapp
tests:
	go test ./...
help:
	@echo "For usage instructions, see README.md."
