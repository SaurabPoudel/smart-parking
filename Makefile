anpr:
	@go build -o bin/anpr anpr/main.go
	@./bin/anpr
receiver:
	@go build -o bin/receiver data_receiver/main.go
	@./bin/receiver

.PHONY: anpr receiver