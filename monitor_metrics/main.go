package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("\n*** Monitoring the block time ***")
		monitorBlockTime()

		time.Sleep(10 * time.Second)
	}
}
