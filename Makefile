anpr:
	@go build -o bin/anpr anpr/main.go
	@./bin/anpr
receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

parking:
	@go build -o bin/parking ./parking
	@./bin/parking
agg:
	@go build -o bin/agg ./aggregator
	@./bin/agg

.PHONY: anpr receiver parking 