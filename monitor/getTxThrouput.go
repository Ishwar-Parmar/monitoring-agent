package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func monitorTransactionThroughput() {
	result, err := sendRPCRequest("eth_blockNumber", []interface{}{})
	if err != nil {
		log.Fatalf("Failed to retrieve block number: %v", err)
	}

	var blockNumber string
	err = json.Unmarshal(result.Result, &blockNumber)
	if err != nil {
		log.Fatalf("Failed to parse block number: %v", err)
	}

	result, err = sendRPCRequest("eth_getBlockByNumber", []interface{}{blockNumber, true})
	if err != nil {
		log.Fatalf("Failed to retrieve block: %v", err)
	}

	var block struct {
		Transactions []interface{} `json:"transactions"`
	}
	err = json.Unmarshal(result.Result, &block)
	if err != nil {
		log.Fatalf("Failed to parse block: %v", err)
	}

	txCount := len(block.Transactions)
	fmt.Printf("Transaction throughput: %d\n", txCount)
}
