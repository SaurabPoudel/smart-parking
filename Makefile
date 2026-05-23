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

proto:
	protoc --go_out=. --go_opt=paths=source_relative types/ptypes.proto

.PHONY: anpr receiver parking 
