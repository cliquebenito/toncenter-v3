package main

import (
	"fmt"
	"github.com/cliquebenito/toncenter-v3"
)

func main() {
	s := toncenter.NewClient("fbb6d4a6b9011e0bc519d01615b588cfc3b1989e1812fbb16d5c813e2ea0be9c")
	// Создание экземпляра Block
	block := &toncenter.Block{
		Client: *s,
	}

	// Вызов метода MasterchainInfo
	info, err := block.MasterchainInfo()
	if err != nil {
		fmt.Println("Error calling MasterchainInfo:", err)
	} else {
		fmt.Println("MasterchainInfo result:", info)
	}

}

// Import resty into your code and refer it as `resty`.
