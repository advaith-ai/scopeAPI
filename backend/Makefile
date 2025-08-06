BIN_DIR=bin
CONFIG_DIR=config

all: api-discovery data-ingestion threat-detection

api-discovery:
	cd services/api-discovery && go build -o ../../$(BIN_DIR)/api-discovery ./cmd

data-ingestion:
	cd services/data-ingestion && go build -o ../../$(BIN_DIR)/data-ingestion ./cmd

threat-detection:
	cd services/threat-detection && go build -o ../../$(BIN_DIR)/threat-detection ./cmd

clean:
	rm -f $(BIN_DIR)/* 