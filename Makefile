ENTRYPOINT = main.go

.PHONY: run
run:
	@HOST=localhost PORT=8080 go run $(ENTRYPOINT)

.PHONY: test
test:
	go test -cover -coverprofile cover.out -covermode=count -v ./...