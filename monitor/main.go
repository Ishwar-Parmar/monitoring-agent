package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("\n*** Monitoring the block time ***")
		monitorBlockTime()

		fmt.Println("\n*** Monitoring the transaction throughput ***")
		monitorTransactionThroughput()

		fmt.Println("\n*** Monitoring the transaction pool status ***")
		monitorTransactionPoolStatus()

		fmt.Println("\n*** Monitoring the gas prices ***")
		monitorGasPrice()

		time.Sleep(10 * time.Second)
	}
}
