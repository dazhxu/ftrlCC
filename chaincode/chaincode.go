package chaincode

import (
	"bsncompetition2/models"
	"bsncompetition2/train"
	"bsncompetition2/utils"

	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type LogisticCC struct {
}

var ftlr *train.FTLR

func (lcc *LogisticCC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	SetLogger("ChainCode Init start......")
	defer SetLogger("ChainCode Init end......")
	SetLogger("初始化FTLR模型...")
	err := initFTLR(stub)
	if err != nil {
		SetLogger("初始化模型错误.")
		return shim.Error(CCResponse(INIT_MODEL_ERR, err.Error()))
	}
	SetLogger("初始化FTLR模型完成")
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

// 设置日志
func SetLogger(logInfo ...interface{}) {
	utils.SetLogger(logInfo)
}

// 生成回应
func CCResponse(code int, msg string) string {
	var ccResp models.Response
	ccResp.Retcode = code
	ccResp.Retmsg = msg
	response, err := json.Marshal(ccResp)
	if err != nil {
		return ""
	}
	return string(response)
}
