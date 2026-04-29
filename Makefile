anpr:
	@go build -o bin/anpr anpr/main.go
	@./bin/anpr
receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

.PHONY: anpr receiver