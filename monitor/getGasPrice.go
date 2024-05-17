package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func monitorGasPrice() {
	result, err := sendRPCRequest("eth_gasPrice", []interface{}{})
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	var gasPrice string
	err = json.Unmarshal(result.Result, &gasPrice)
	if err != nil {
		log.Fatalf("Failed to parse gas price: %v", err)
	}

	fmt.Printf("Gas price: %s\n", gasPrice)
}
