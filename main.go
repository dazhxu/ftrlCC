package main

import (
	"bsncompetition2/chaincode"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(chaincode.LogisticCC))
	if err != nil {
		fmt.Printf("Error starting BsnChainCode: %s", err)
	}
}
