### Required tools
GOTOOLS_CHECK = dep gin golangci-lint

### Testing
test:
	@echo "==> Running tests..."
	go test ./... -covermode=atomic -coverpkg=./... -count=1 -race

### Formatting and linting
fmt:
	@echo "==> Running fmt..."
	go fmt ./...

linter:
	@echo "==> Running linter..."
	golangci-lint run ./...

# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: test fmt linter
