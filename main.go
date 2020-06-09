package main

import (
	"fmt"
	"ftrlCC/chaincode"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {
	err := shim.Start(new(chaincode.LogisticCC))
	if err != nil {
		fmt.Printf("Error starting BsnChainCode: %s", err)
	}
}
