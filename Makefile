anpr:
	@go build -o bin/anpr anpr/main.go
	@./bin/anpr
receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

parking:
	@go build -o bin/parking ./parking
	@./bin/parking

.PHONY: anpr receiver parking