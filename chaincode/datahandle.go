package chaincode

import (
	"bsncompetition2/models"
	"bsncompetition2/train"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/msp"
)

// 添加记录
func recordCount(stub shim.ChaincodeStubInterface, num int) error {
	creatorByte, err := stub.GetCreator()
	fmt.Println("Creator: " + string(creatorByte))

	si := &msp.SerializedIdentity{}

	err = proto.Unmarshal(creatorByte, si)
	if err != nil {
		return errors.New("Cannot get identity of creator.")
	}
	mspid := si.Mspid

	// 尝试从state获取记录数据
	recordsBytes, err := stub.GetState(TRAIN_RECORDS_KEY)
	if err != nil {
		return errors.New("Cannot get state.")
	}
	// 添加记录
	var records []models.Record
	if recordsBytes == nil {
		records = append(records, models.Record{mspid, num})
	} else {
		err = json.Unmarshal(recordsBytes, &records)
		if err != nil {
			return errors.New("Cannot unmarshal record str to records struct.")
		}
		exist := false
		for idx, item := range records {
			if item.MspId == mspid {
				records[idx].Count += num
				exist = true
			}
		}
		if !exist {
			records = append(records, models.Record{mspid, num})
		}
	}

	// 将记录写回state
	recordsBytes, err = json.Marshal(records)
	if err != nil {
		return errors.New("Cannot marshal records to bytes.")
	}

	err = stub.PutState(TRAIN_RECORDS_KEY, recordsBytes)
	if err != nil {
		return errors.New("Cannot put records bytes to state.")
	}
	return nil
}

// 初始化FTLR
func initFTLR(stub shim.ChaincodeStubInterface) error {
	if ftlr != nil {
		return nil
	}

	ftlrBytes, err := stub.GetState(FTLR_MODEL_KEY)
	if err != nil {
		return err
	}
	if ftlrBytes == nil {
		ftlr = train.Init(4, 1.0, 1.0, 0.1, 1.0, new(train.LR))
	} else {
		err = json.Unmarshal(ftlrBytes, ftlr)
		if err != nil {
			ftlr = train.Init(4, 1.0, 1.0, 0.1, 1.0, new(train.LR))
		}
	}
	ftlrBytes, err = json.Marshal(ftlr)
	if err != nil {
		return err
	}
	err = stub.PutState(FTLR_MODEL_KEY, ftlrBytes)
	if err != nil {
		return err
	}
	return nil

}
