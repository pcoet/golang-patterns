.PHONY: clean-bin build test help info
.DEFAULT_GOAL := help

UHOME = ${shell echo "${HOME}"}
GOPATH = ${UHOME}/go
GOBIN = ${GOPATH}/bin
APP_NAME = testapp
INSTALLATION = ${GOBIN}/${APP_NAME}
MODULE = github.com/pcoet/golang-patterns
PACKAGE = ${MODULE}/cmd/${APP_NAME}

clean-bin:
	@echo "Deleting installed binary: ${INSTALLATION}"
	@rm -rf ${INSTALLATION}
build: clean-bin
	@echo "Installing ${PACKAGE} to ${GOBIN}"
	@go install ${PACKAGE}
test:
	go test ./...
help:
	@echo "For usage instructions, see README.md."
info:
	@echo "UHOME:        ${UHOME}"
	@echo "GOPATH:       ${GOPATH}"
	@echo "GOBIN:        ${GOBIN}"
	@echo "APP_NAME:     ${APP_NAME}"
	@echo "INSTALLATION: ${INSTALLATION}"
	@echo "MODULE:       ${MODULE}"
	@echo "PACKAGE:      ${PACKAGE}"
