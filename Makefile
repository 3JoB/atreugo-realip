.PHONY: tidy
tidy:
	go mod tidy

.PHONY: test
test:
	gotestsum -- -race -p=1 ./...

fmt:
	goimports -w .

PHONY: dep
dep:
	go install golang.org/x/tools/cmd/goimports@latest

	$(MAKE) tidy