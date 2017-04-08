package main

import (
	
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	

)
type SN struct {

	cl1	CL
}


type SNJSON struct {

		
		RefNo 		string `json:"RefNo"`
		ExporterName string `json:"ExporterName"`
		ImporterName string `json:"ImporterName"`
		Commodity string `json:"Commodity"`
		UnitPrice string `json:"UnitPrice"`
		Amount string `json:"Amount"`
		Currency string `json:"Currency"`
		Quantity string `json:"Quantity"`
		Weight 	string `json:"Weight"`
		TermsOfTrade string `json:"TermsOfTrade"`
		TermsOfInsurance string `json:"TermsOfInsurance"`
		TermsOfPayment string `json:"TermsOfPayment"`
		PackingMethod	string `json:"PackingMethod"`
		WayOfTransportation string `json:"WayOfTransportation"`
		TimeOfShipment string `json:"TimeOfShipment"`
		PortOfShipment string `json:"PortOfShipment"`
		PortOfDischarge string `json:"PortOfDischarge"`
		ProcessStatus   string `json:"ProcessStatus"`
		SNSubmittedTime string `json: "SNSubmittedTime"`
		SNRejectReason string `json:"SNRejectReason"`
		PaymentDate string `json:"PaymentDate"`


}
	

	type SNContractsList struct {
		ContractNo string `json:"ContractNo"`
	}

	type count struct {
		NumContracts int
	}

	type ListContracts struct{
		snDetail	[]ListSN `json:"listSN"`
	}

	type ListSN struct {

		ContractNo 		string `json:"ContractNo"`
		SNSubmittedTime string `json: "SNSubmittedTime"`
		ImporterName string `json:"ImporterName"`
		ExporterName string `json:"ExporterName"`
		Currency string `json:"Currency"`
		Amount string `json:"Amount"`
		Commodity string `json:"Commodity"`
		ProcessStatus   string `json:"ProcessStatus"`
		CargoLocation string `json:"CargoLocation"`	
		
			}

	type SNList struct {

			contractNo  []ContractNo `json:"contractNo"`
		}

	type ContractNo struct {

		ContractNo  string `json:"ContractNo"`
	}


	type Cl1Args struct {

		ContractNo 		string `json:"ContractNo"`
		CargoLocation 	string `json:"CargoLocation"`

	}  

func (t *SN) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Check if table already exists
	_, err := stub.GetTable("SalesNote")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create SN Table
	err = stub.CreateTable("SalesNote", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "Type", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "ContractNo", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "RefNo", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ExporterName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ImporterName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Commodity", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "UnitPrice", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Amount", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Currency", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Quantity", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "Weight", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "TermsOfTrade", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "TermsOfInsurance", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "TermsOfPayment", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PackingMethod", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "WayOfTransportation", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "TimeOfShipment", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PortOfShipment", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PortOfDischarge", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "ProcessStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "SNInitialCreateTime", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "UpdateTime", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "SNCreatedTime", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "SNSubmittedTime", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "CompanyIdOfExporter", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "CompanyIdOfImporter", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "SNRejectReason", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "PaymentDate", Type: shim.ColumnDefinition_STRING, Key: false},

	})
	if err != nil {
		return nil, errors.New("Failed creating SalesNoteTable.")
	}

	return nil, nil

}


//SubmitDoc () inserts a new row in the table

func (t *SN) SubmitDoc(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

		
		if len(args) != 25 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 25. Got: %d.", len(args))
		}

		ContractNo := args[0]
		RefNo := args[1]
		ExporterName := args[2]
		ImporterName := args[3]
		Commodity := args[4]
		UnitPrice := args[5]
		Amount := args[6]
		Currency := args[7]
		Quantity:= args[8]
		Weight := args[9]
		TermsOfTrade := args[10]
		TermsOfInsurance := args[11]
		TermsOfPayment := args[12]
		PackingMethod:= args[13]
		WayOfTransportation := args[14]
		TimeOfShipment := args[15]
		PortOfShipment := args[16]
		PortOfDischarge := args[17]
		//ProcessStatus := args[18]
		SNInitialCreateTime := args[18]
		UpdateTime := args[19]
		SNCreatedTime := args[20]
		SNSubmittedTime := args[21]
		CompanyIdOfExporter := args[22]
		CompanyIdOfImporter := args[23]
		SNRejectReason := ""
		PaymentDate := args[24]



		// Insert a row
	ok, err := stub.InsertRow("SalesNote", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "SN"}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractNo}},
			&shim.Column{Value: &shim.Column_String_{String_: RefNo}},
			&shim.Column{Value: &shim.Column_String_{String_: ExporterName}},
			&shim.Column{Value: &shim.Column_String_{String_: ImporterName}},
			&shim.Column{Value: &shim.Column_String_{String_: Commodity}},
			&shim.Column{Value: &shim.Column_String_{String_: UnitPrice}},
			&shim.Column{Value: &shim.Column_String_{String_: Amount}},
			&shim.Column{Value: &shim.Column_String_{String_: Currency}},
			&shim.Column{Value: &shim.Column_String_{String_: Quantity}},
			&shim.Column{Value: &shim.Column_String_{String_: Weight}},
			&shim.Column{Value: &shim.Column_String_{String_: TermsOfTrade}},
			&shim.Column{Value: &shim.Column_String_{String_: TermsOfInsurance}},
			&shim.Column{Value: &shim.Column_String_{String_: TermsOfPayment}},
			&shim.Column{Value: &shim.Column_String_{String_: PackingMethod}},
			&shim.Column{Value: &shim.Column_String_{String_: WayOfTransportation}},
			&shim.Column{Value: &shim.Column_String_{String_: TimeOfShipment}},
			&shim.Column{Value: &shim.Column_String_{String_: PortOfShipment}},
			&shim.Column{Value: &shim.Column_String_{String_: PortOfDischarge}},
			&shim.Column{Value: &shim.Column_String_{String_: "S/N Created"}},
			&shim.Column{Value: &shim.Column_String_{String_: SNInitialCreateTime}},
			&shim.Column{Value: &shim.Column_String_{String_: UpdateTime}},
			&shim.Column{Value: &shim.Column_String_{String_: SNCreatedTime}},
			&shim.Column{Value: &shim.Column_String_{String_: SNSubmittedTime}},
			&shim.Column{Value: &shim.Column_String_{String_: CompanyIdOfExporter}},
			&shim.Column{Value: &shim.Column_String_{String_: CompanyIdOfImporter}},
			&shim.Column{Value: &shim.Column_String_{String_: SNRejectReason}},
			&shim.Column{Value: &shim.Column_String_{String_: PaymentDate}}},

	})

	if !ok && err == nil {
		return nil, errors.New("Document already exists.")
	}


			toSend := make ([]string, 2)
			toSend[0] = string(ContractNo)
			toSend[1] = ""
			
			_,cl1Err := t.cl1.UpdateCargoLocation(stub, toSend)
			if cl1Err != nil {
				return nil, cl1Err
			} 

	
	return nil, err
	}



	func (t *SN) UpdateSN(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){

		ContractNo := args[0]
		rejectionComment := "" 
		snStatus := ""

		if len(args) == 3 {
			snStatus = args[2]
			rejectionComment = args[1]

		} else if len(args) == 2 {

		snStatus = args[1]
		

		} else {

			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 2 or 3. Got: %d.", len(args))
		}

		// Get the row pertaining to this UID
		var columns []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: "SN"}}
		columns = append(columns, col1)
		col2 := shim.Column{Value: &shim.Column_String_{String_: ContractNo}}
		columns = append(columns, col2)

		row, err := stub.GetRow("SalesNote", columns)
		if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving document with ContractNo %s. Error %s", ContractNo, err.Error())
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
		return nil, err
		}


		
		RefNo := row.Columns[2].GetString_()
		ExporterName := row.Columns[3].GetString_()
		ImporterName := row.Columns[4].GetString_()
		Commodity := row.Columns[5].GetString_()
		UnitPrice := row.Columns[6].GetString_()
		Amount := row.Columns[7].GetString_()
		Currency := row.Columns[8].GetString_()
		Quantity:= row.Columns[9].GetString_()
		Weight := row.Columns[10].GetString_()
		TermsOfTrade := row.Columns[11].GetString_()
		TermsOfInsurance := row.Columns[12].GetString_()
		TermsOfPayment := row.Columns[13].GetString_()
		PackingMethod:= row.Columns[14].GetString_()
		WayOfTransportation := row.Columns[15].GetString_()
		TimeOfShipment := row.Columns[16].GetString_()
		PortOfShipment := row.Columns[17].GetString_()
		PortOfDischarge := row.Columns[18].GetString_()
		ProcessStatus := row.Columns[19].GetString_()
		SNInitialCreateTime := row.Columns[20].GetString_()
		UpdateTime := row.Columns[21].GetString_()
		SNCreatedTime := row.Columns[22].GetString_()
		SNSubmittedTime := row.Columns[23].GetString_()
		CompanyIdOfExporter := row.Columns[24].GetString_()
		CompanyIdOfImporter := row.Columns[25].GetString_()
		SNRejectReason := rejectionComment
		PaymentDate := row.Columns[27].GetString_()




		var newStatus string

		if snStatus == "AcceptSN"{

			newStatus = "S/N Submitted"

		} else if snStatus == "RejectSN"{

			newStatus = "S/N Rejected"
		}

		//Start- Check that the currentStatus to newStatus transition is accurate

		stateTransitionAllowed := false

		if ProcessStatus == "S/N Created" && newStatus == "S/N Submitted" {
		stateTransitionAllowed = true
		} else if ProcessStatus == "S/N Created" && newStatus == "S/N Rejected" {
		stateTransitionAllowed = true
		} 

	if stateTransitionAllowed == false {
		return nil, errors.New("This state transition is not allowed.")
	}

	//End- Check that the currentStatus to newStatus transition is accurate

		ok, err := stub.ReplaceRow("SalesNote", shim.Row{
		Columns: []*shim.Column{
			&shim.Column{Value: &shim.Column_String_{String_: "SN"}},
			&shim.Column{Value: &shim.Column_String_{String_: ContractNo}},
			&shim.Column{Value: &shim.Column_String_{String_: RefNo}},
			&shim.Column{Value: &shim.Column_String_{String_: ExporterName}},
			&shim.Column{Value: &shim.Column_String_{String_: ImporterName}},
			&shim.Column{Value: &shim.Column_String_{String_: Commodity}},
			&shim.Column{Value: &shim.Column_String_{String_: UnitPrice}},
			&shim.Column{Value: &shim.Column_String_{String_: Amount}},
			&shim.Column{Value: &shim.Column_String_{String_: Currency}},
			&shim.Column{Value: &shim.Column_String_{String_: Quantity}},
			&shim.Column{Value: &shim.Column_String_{String_: Weight}},
			&shim.Column{Value: &shim.Column_String_{String_: TermsOfTrade}},
			&shim.Column{Value: &shim.Column_String_{String_: TermsOfInsurance}},
			&shim.Column{Value: &shim.Column_String_{String_: TermsOfPayment}},
			&shim.Column{Value: &shim.Column_String_{String_: PackingMethod}},
			&shim.Column{Value: &shim.Column_String_{String_: WayOfTransportation}},
			&shim.Column{Value: &shim.Column_String_{String_: TimeOfShipment}},
			&shim.Column{Value: &shim.Column_String_{String_: PortOfShipment}},
			&shim.Column{Value: &shim.Column_String_{String_: PortOfDischarge}},
			&shim.Column{Value: &shim.Column_String_{String_: newStatus}},
			&shim.Column{Value: &shim.Column_String_{String_: SNInitialCreateTime}},
			&shim.Column{Value: &shim.Column_String_{String_: UpdateTime}},
			&shim.Column{Value: &shim.Column_String_{String_: SNCreatedTime}},
			&shim.Column{Value: &shim.Column_String_{String_: SNSubmittedTime}},
			&shim.Column{Value: &shim.Column_String_{String_: CompanyIdOfExporter}},
			&shim.Column{Value: &shim.Column_String_{String_: CompanyIdOfImporter}},
			&shim.Column{Value: &shim.Column_String_{String_: SNRejectReason}},
			&shim.Column{Value: &shim.Column_String_{String_: PaymentDate}},
			
			
		}})

	if !ok && err == nil {

		return nil, errors.New("Document unable to Update.")
	}



		if snStatus == "AcceptSN"{
			toSend := make ([]string, 2)
			toSend[0] = string(ContractNo)
			toSend[1] = "Exporter"

			
			_,cl1Err := t.cl1.UpdateCargoLocation(stub, toSend)
			if cl1Err != nil {
				return nil, cl1Err
			} 
		}

	
		return nil, nil
}


	func (t *SN) GetContractStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){
		if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
		}


		var snJSON SNJSON

		sndetails,_ := t.GetSN(stub, []string{args[0]})

		err := json.Unmarshal(sndetails, &snJSON)
		if err != nil{

			return nil, err

		}

		
		return []byte(snJSON.ProcessStatus), nil

	}


	func (t *SN) GetSN (stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

		if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
		}

	


	ContractNo := args[0]

	// Get the row pertaining to this UID
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: "SN"}}
	columns = append(columns, col1)
	col2 := shim.Column{Value: &shim.Column_String_{String_: ContractNo}}
	columns = append(columns, col2)

	row, err := stub.GetRow("SalesNote", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving document with ContractNo %s. Error %s", ContractNo, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}

	var snJSON SNJSON 

	snJSON.RefNo = row.Columns[2].GetString_()
	snJSON.ExporterName = row.Columns[3].GetString_()
	snJSON.ImporterName = row.Columns[4].GetString_()
	snJSON.Commodity = row.Columns[5].GetString_() 
	snJSON.UnitPrice = row.Columns[6].GetString_()
	snJSON.Amount = row.Columns[7].GetString_()
	snJSON.Currency = row.Columns[8].GetString_()
	snJSON.Quantity = row.Columns[9].GetString_()
	snJSON.Weight = row.Columns[10].GetString_()
	snJSON.TermsOfTrade = row.Columns[11].GetString_()
	snJSON.TermsOfInsurance = row.Columns[12].GetString_()
	snJSON.TermsOfPayment = row.Columns[13].GetString_()
	snJSON.PackingMethod = row.Columns[14].GetString_()
	snJSON.WayOfTransportation = row.Columns[15].GetString_()
	snJSON.TimeOfShipment = row.Columns[16].GetString_()
	snJSON.PortOfShipment = row.Columns[17].GetString_()
	snJSON.PortOfDischarge = row.Columns[18].GetString_()
	snJSON.ProcessStatus = row.Columns[19].GetString_()
	snJSON.SNSubmittedTime = row.Columns[23].GetString_()
	snJSON.SNRejectReason = row.Columns[26].GetString_()
	snJSON.PaymentDate = row.Columns[27].GetString_()
	

	jsonSN, err := json.Marshal(snJSON)

	if err != nil {

		return nil, err
	}

	fmt.Println(jsonSN)

 	return jsonSN, nil

	}



	func (t *SN) ListContractsByCompID(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

		if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
		}



		companyID := args[0]
		roleID := args[1]
	
		var listContracts ListContracts

		listContracts.snDetail = make([]ListSN, 0)
		

		if roleID == "1" || roleID == "4" {

			
			var columns []shim.Column
			col1 := shim.Column{Value: &shim.Column_String_{String_: "SN"}}
			columns = append(columns, col1)
			

			rows,err := stub.GetRows("SalesNote", columns)
			
			if err != nil {
			return nil, fmt.Errorf("Error: Failed retrieving document Error %s", err.Error())
			}

			
			
			var contractIDOfUser ContractsList

			

				for row := range rows {
					
					contractIDOfUser.ContractNo = ""
				
					if len(row.Columns) == 0 { 

						break 
					
					} else if roleID == "1" {
						if row.Columns[24].GetString_() == companyID{

					 		contractIDOfUser.ContractNo = row.Columns[1].GetString_()

						}
					} else if roleID == "4"{

						if row.Columns[25].GetString_() == companyID{

					 	contractIDOfUser.ContractNo = row.Columns[1].GetString_()

						}

					}  


					if contractIDOfUser.ContractNo != "" {

					b,_ := t.GetSN(stub, []string{contractIDOfUser.ContractNo})
					var listSN ListSN
					err = json.Unmarshal(b, &listSN)
					listSN.ContractNo = contractIDOfUser.ContractNo

					if listSN.ProcessStatus == "S/N Rejected" {

						listSN.ProcessStatus = ""
					}

					b,_ = t.cl.GetCargoLocation(stub, []string{contractIDOfUser.ContractNo})

					if err != nil {
        				return nil, err
        			}
					listSN.CargoLocation = string(b)
					listContracts.snDetail = append(listContracts.snDetail, listSN)

					}

				}

				

		} else if roleID == "2" || roleID == "3" || roleID == "6" {

			

			var columns []shim.Column
			col1 := shim.Column{Value: &shim.Column_String_{String_: "RR"}}
			columns = append(columns, col1)
			

			rows,err := stub.GetRows("RequestReservation", columns)
			
			if err != nil {
			return nil, fmt.Errorf("Error: Failed retrieving document Error %s", err.Error())
			
			}


			var contractIDOfUser ContractsList
			
				for row := range rows {
	
					contractIDOfUser.ContractNo = ""
					if len(row.Columns) == 0 {

						break 
					
					} else if roleID == "2" {
						if row.Columns[7].GetString_() == companyID{

					 		contractIDOfUser.ContractNo = row.Columns[1].GetString_()

					 		
						}
					} else if roleID == "3" {


						if row.Columns[8].GetString_() == companyID{

					 	contractIDOfUser.ContractNo = row.Columns[1].GetString_()

					 	
						}

					} else if roleID == "6"{

						if row.Columns[4].GetString_() == companyID{

					 	contractIDOfUser.ContractNo = row.Columns[1].GetString_()

					 	
						}

					} 

					if contractIDOfUser.ContractNo != "" {

					//var snJSON SNJSON
					b,_ := t.GetSN(stub, []string{contractIDOfUser.ContractNo})
					var listSN ListSN
					err = json.Unmarshal(b, & listSN)
					listSN.ContractNo = contractIDOfUser.ContractNo

					if listSN.ProcessStatus == "S/N Rejected" {

						listSN.ProcessStatus = ""
					}

					b,_ = t.cl.GetCargoLocation(stub, []string{contractIDOfUser.ContractNo})
				
					listSN.CargoLocation = string(b)

					listContracts.snDetail = append(listContracts.snDetail, listSN)

				}

				}

		} else if roleID == "5" {

			var columns []shim.Column
			col1 := shim.Column{Value: &shim.Column_String_{String_: "CRR"}}
			columns = append(columns, col1)
			

			rows,err := stub.GetRows("CargoReceiveRequest", columns)
			
			if err != nil {
			return nil, fmt.Errorf("Error: Failed retrieving document Error %s", err.Error())
			}

			var contractIDOfUser ContractsList
			

				for row := range rows {

					contractIDOfUser.ContractNo = ""

					if len(row.Columns) == 0 {
					
						
						break 
					
					} else if roleID == "5" {
						if row.Columns[5].GetString_() == companyID{

					 		contractIDOfUser.ContractNo = row.Columns[1].GetString_()

					 		
						}
					}  

					if contractIDOfUser.ContractNo != "" {
					b,_ := t.GetSN(stub, []string{contractIDOfUser.ContractNo})
					var listSN ListSN
					err = json.Unmarshal(b, &listSN)
					listSN.ContractNo = contractIDOfUser.ContractNo


					if listSN.ProcessStatus == "S/N Rejected" {

						listSN.ProcessStatus = ""
					}


					b,_ = t.cl.GetCargoLocation(stub, []string{contractIDOfUser.ContractNo})

					listSN.CargoLocation = string(b)

					listContracts.snDetail = append(listContracts.snDetail, listSN)

					}

				}
			
		}

				

		return json.Marshal(listContracts.snDetail) 

	}


	func (t *SN) NumbContracts(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

		if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0.")
	}

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: "SN"}}
	columns = append(columns, col1)

	contractCounter := 0

	rows, err := stub.GetRows("SalesNote", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	for row := range rows {
		if len(row.Columns) != 0 {
			contractCounter++
		}
	}

	var c count
	c.NumContracts = contractCounter

	return json.Marshal(c)


	}


	func (t *SN) ListOfContracts(stub shim.ChaincodeStubInterface, args []string) ([]byte, error){
		if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0.")
		}

	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: "SN"}}
	columns = append(columns, col1)
	

	rows, err := stub.GetRows("SalesNote", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

 	var listOfContracts SNList
 	var contractNo ContractNo 

 	listOfContracts.contractNo = make([]ContractNo, 0)


	for row := range rows {

		
		 contractNo.ContractNo = row.Columns[1].GetString_()

        listOfContracts.contractNo = append(listOfContracts.contractNo, contractNo)

	}

		return json.Marshal(listOfContracts.contractNo)

	}
