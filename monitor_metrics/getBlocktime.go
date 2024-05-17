package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func monitorBlockTime() {
	result, err := sendRPCRequest("eth_blockNumber", []interface{}{})
	if err != nil {
		log.Fatalf("Failed to retrieve block number: %v", err)
	}

	var blockNumber string
	err = json.Unmarshal(result.Result, &blockNumber)
	if err != nil {
		log.Fatalf("Failed to parse block number: %v", err)
	}

	fmt.Println("Latest block number: ", blockNumber)
}
