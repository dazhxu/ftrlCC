package chaincode

import (
	"bsncompetition2/models"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const (
	TRAIN_RECORDS_KEY = "TRAIN_RECORDS"
	SUCCESS           = 0
	PARAM_ERR         = 1001
	MARSHAL_ERR       = 1002
	STATE_ERR         = 1003
	RECORD_ERR        = 1004
)

// 使用一条数据进行训练
func trainOnce(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error(CCResponse(PARAM_ERR, "The length of args should be 1"))
	}
	// 数据预处理
	dataJson := args[0]
	var dataEntry models.DataEntry
	err := json.Unmarshal([]byte(dataJson), &dataEntry)
	if err != nil {
		return shim.Error(CCResponse(MARSHAL_ERR, fmt.Sprintf("Cannot unmarshal %s to DataEntry", dataJson)))
	}

	// 添加记录
	err = recordCount(stub, 1)
	if err != nil {
		return shim.Error(CCResponse(RECORD_ERR, err.Error()))
	}

	// TODO 训练逻辑

	return shim.Success(nil)
}

// 使用一批数据进行训练
func trainBatch(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error(CCResponse(PARAM_ERR, "The length of args should be 1"))
	}
	// 数据预处理
	dataJson := args[0]
	var dataEntries []models.DataEntry
	err := json.Unmarshal([]byte(dataJson), &dataEntries)
	if err != nil {
		return shim.Error(CCResponse(MARSHAL_ERR, fmt.Sprintf("Cannot unmarshal %s to DataEntry", dataJson)))
	}

	// 添加记录
	err = recordCount(stub, len(dataEntries))
	if err != nil {
		return shim.Error(CCResponse(RECORD_ERR, err.Error()))
	}

	// TODO 训练逻辑

	return shim.Success(nil)
}

// 预测
func predict(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error(CCResponse(PARAM_ERR, "The length of args should be 1"))
	}
	// 数据预处理
	dataJson := args[0]
	var dataEntry models.DataEntry
	err := json.Unmarshal([]byte(dataJson), &dataEntry)
	if err != nil {
		return shim.Error(CCResponse(MARSHAL_ERR, fmt.Sprintf("Cannot unmarshal %s to DataEntry", dataJson)))
	}

	// TODO 预测逻辑

	return shim.Success(nil)
}

// 输出各方数据统计结果
func statistics(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	recordsBytes, err := stub.GetState(TRAIN_RECORDS_KEY)
	if err != nil || recordsBytes == nil {
		return shim.Error(CCResponse(STATE_ERR, "Cannot get records from state."))
	}

	return shim.Success([]byte(CCResponse(SUCCESS, "SUCCESS")))
}
