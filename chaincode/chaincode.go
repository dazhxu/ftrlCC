package chaincode

import (
	"bsncompetition2/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type LogisticCC struct {
}

// 设置日志
func SetLogger(logInfo ...interface{}) {
	utils.SetLogger(logInfo)
}

func (lcc *LogisticCC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	SetLogger("ChainCode Init start......")
	defer SetLogger("ChainCode Init end......")
	return shim.Success(nil)
}

func (lcc *LogisticCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "trainOnce":
		return trainOnce(stub, args)
	case "trainBatch":
		return trainBatch(stub, args)
	case "predict":
		return predict(stub, args)
	case "statistics":
		return statistics(stub, args)
	default:
		SetLogger("无效的方法")
		return shim.Error("无效的请求")
	}
}
