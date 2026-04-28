anpr:
	@go build -o bin/anpr anpr/main.go
	@./bin/anpr

.PHONY: anpr