package chaincode

import (
	"encoding/json"
	"fmt"
	"ftrlCC/models"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const (
	TRAIN_RECORDS_KEY = "TRAIN_RECORDS"
	FTRL_MODEL_KEY    = "FTRL_MODEL"
	//LOSS_KEY 		  = "LOSS"
	SUCCESS        = 0
	INIT_MODEL_ERR = 1000
	PARAM_ERR      = 1001
	MARSHAL_ERR    = 1002
	STATE_ERR      = 1003
	RECORD_ERR     = 1004
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

	// 训练逻辑
	if ftrl == nil {
		initFTRL(stub)
	}

	loss := ftrl.Update(dataEntry.X, dataEntry.Y)

	// 将FTRL模型写入state
	ftrlBytes, err := json.Marshal(ftrl)
	if err != nil {
		return shim.Error(CCResponse(MARSHAL_ERR, "Cannot marshal ftrl model."))
	}
	err = stub.PutState(FTRL_MODEL_KEY, ftrlBytes)
	if err != nil {
		return shim.Error(CCResponse(STATE_ERR, "Cannot put ftrl model to state."))
	}

	return shim.Success([]byte(CCResponse(SUCCESS, fmt.Sprintf("Final loss is %.6f", loss))))
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

	// 训练逻辑
	if ftrl == nil {
		initFTRL(stub)
	}

	var loss float64
	for _, entry := range dataEntries {
		loss = ftrl.Update(entry.X, entry.Y)
	}

	// 将FTRL模型写入state
	ftrlBytes, err := json.Marshal(ftrl)
	if err != nil {
		return shim.Error(CCResponse(MARSHAL_ERR, "Cannot marshal ftrl model."))
	}
	err = stub.PutState(FTRL_MODEL_KEY, ftrlBytes)
	if err != nil {
		return shim.Error(CCResponse(STATE_ERR, "Cannot put ftrl model to state."))
	}

	return shim.Success([]byte(CCResponse(SUCCESS, fmt.Sprintf("Final loss is %.6f", loss))))
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

	// 预测逻辑
	if ftrl == nil {
		initFTRL(stub)
	}

	dataEntry.Y = ftrl.Predict(dataEntry.X)
	dataEntryBytes, _ := json.Marshal(dataEntry)

	return shim.Success([]byte(CCResponse(SUCCESS, string(dataEntryBytes))))
}

// 输出各方数据统计结果
func statistics(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	recordsBytes, err := stub.GetState(TRAIN_RECORDS_KEY)
	if err != nil || recordsBytes == nil {
		return shim.Error(CCResponse(STATE_ERR, "Cannot get records from state."))
	}

	return shim.Success([]byte(CCResponse(SUCCESS, string(recordsBytes))))
}
