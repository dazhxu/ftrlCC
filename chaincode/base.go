package chaincode

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// 使用一条数据进行训练
func trainOnce(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success(nil)
}

// 使用一批数据进行训练
func trainBatch(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success(nil)
}

// 预测
func predict(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success(nil)
}

// 输出各方数据统计结果
func statistics(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success(nil)
}
