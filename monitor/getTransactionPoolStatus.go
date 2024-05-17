package main

import (
	"fmt"
	"log"
)

func monitorTransactionPoolStatus() {
	result, err := sendRPCRequest("txpool_status", []interface{}{})
	if err != nil {
		log.Fatalf("Failed to get monitor transaction pool status: %v", err)
	}

	fmt.Println(result)
}
