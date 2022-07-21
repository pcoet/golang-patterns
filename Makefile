.PHONY: clean-bin build run test help
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
test:
	go test ./...
help:
	@echo "For usage instructions, see README.md."
