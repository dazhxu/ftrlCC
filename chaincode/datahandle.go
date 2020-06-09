package chaincode

import (
	"encoding/json"
	"errors"
	"ftrlCC/models"
	"ftrlCC/train"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// 添加记录
func recordCount(stub shim.ChaincodeStubInterface, num int) error {
	creatorByte, err := stub.GetCreator()
	//fmt.Println("Creator: " + string(creatorByte))

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

// 初始化FTRL
func initFTRL(stub shim.ChaincodeStubInterface) error {
	if ftrl != nil {
		return nil
	}

	ftrlBytes, err := stub.GetState(FTRL_MODEL_KEY)
	if err != nil {
		return err
	}
	if ftrlBytes == nil {
		ftrl = train.Init(4, 1.0, 1.0, 0.1, 1.0, new(train.LR))
	} else {
		err = json.Unmarshal(ftrlBytes, ftrl)
		if err != nil {
			ftrl = train.Init(4, 1.0, 1.0, 0.1, 1.0, new(train.LR))
		}
	}
	ftrlBytes, err = json.Marshal(ftrl)
	if err != nil {
		return err
	}
	err = stub.PutState(FTRL_MODEL_KEY, ftrlBytes)
	if err != nil {
		return err
	}
	return nil

}
