package main

import (
	
	"errors"
	//"fmt"
	//"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)
type BC struct {

}

func (t *BC) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Check if table already exists
	_, err := stub.GetTable("BookingConfirmation")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}


	// Create BC Table
	err = stub.CreateTable("BookingConfirmation", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ContractNo", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ContainerNo", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PlaceOfDelivery", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "NumberOfContainers", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PlaceOfReceipt", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ExpectedTimeOfDepature", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ExpectedTimeOfArrival", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "CutOffDateTime", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "VesselName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "VesselNo", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "BookingNo", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Freight", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "FreightPayment", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "UpdateTime", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "BCSubmittedTime", Type: shim.ColumnDefinition_STRING, Key: false},


	})
	if err != nil {
		return nil, errors.New("Failed creating BookingConfirmationTable.")
	}


	return nil, nil

	}
