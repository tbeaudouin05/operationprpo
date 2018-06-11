package baainteract

import (
	"database/sql"
	"log"

	"github.com/thomas-bamilo/operation/operationprpo/row/costcenter"

	"github.com/thomas-bamilo/operation/operationprpo/row/purchaserequestforminput"
	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"
)

// CHANGE TO ID COST CENTER

// LoadPurchaseRequestToDb loads purchase request to baa_application.operation.purchase_request
func LoadPurchaseRequestToDb(purchaseRequestFormInput *purchaserequestforminput.PurchaseRequestFormInput, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertPurchaseRequestStr := `
INSERT INTO baa_application.operation.purchase_request (
	gfk_cost_center  
	,initiator  
	,pr_type  
	,cost_category  
	,invoice_number  
	,invoice_date  
	,vendor_name  
	,item_description  
	,unit_price  
	,vat_unit_price  
	,quantity  
	,payment_term  
	,payment_installment
	,payment_center   
	,payment_type
	,purchase_request_status) 
VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13,@p14,@p15,'pending')`
	insertPurchaseRequest, err := dbBaa.Prepare(insertPurchaseRequestStr)

	res, err := insertPurchaseRequest.Exec(
		purchaseRequestFormInput.CostCenter,
		purchaseRequestFormInput.Initiator,
		purchaseRequestFormInput.PrType,
		purchaseRequestFormInput.CostCategory,
		purchaseRequestFormInput.InvoiceNumber,
		purchaseRequestFormInput.InvoiceDate,
		purchaseRequestFormInput.VendorName,
		purchaseRequestFormInput.ItemDescription,
		purchaseRequestFormInput.UnitPrice,
		purchaseRequestFormInput.VatUnitPrice,
		purchaseRequestFormInput.Quantity,
		purchaseRequestFormInput.PaymentTerm,
		purchaseRequestFormInput.PaymentInstallment,
		purchaseRequestFormInput.PaymentCenter,
		purchaseRequestFormInput.PaymentType,
	)

	log.Println(res)

	return err
}

// LoadPurchaseRequestToDb loads purchase request to baa_application.operation.purchase_request
func CreateNewUser(userFormInput *useraccess.User, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertNewUserStr := `
INSERT INTO baa_application.operation.pr_user (
	email 
	,name
	,access) 
VALUES (@p1,@p2,'user')`
	insertNewUser, err := dbBaa.Prepare(insertNewUserStr)

	res, err := insertNewUser.Exec(
		userFormInput.Email,
		userFormInput.Name,
	)

	log.Println(res)

	return err
}

func CreateNewCostCenter(costCenterFormInput *costcenter.CostCenter, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertNewCostCenterStr := `
	INSERT INTO baa_application.operation.func (
		gfk_department
		,function_code
		,name
		) VALUES (@p1,@p2,@p3)`
	insertNewCostCenter, err := dbBaa.Prepare(insertNewCostCenterStr)

	res, err := insertNewCostCenter.Exec(
		costCenterFormInput.GFKDepartment,
		costCenterFormInput.FunctionCode,
		costCenterFormInput.FunctionName,
	)

	log.Println(res)

	return err
}

func CreateNewDepartment(departmentFormInput *costcenter.Department, dbBaa *sql.DB) error {

	insertNewDepartmentStr := `
	INSERT INTO baa_application.operation.department (
		gfk_location
		,department_code
		,name
		) VALUES (@p1,@p2,@p3)`
	insertNewDepartment, err := dbBaa.Prepare(insertNewDepartmentStr)

	res, err := insertNewDepartment.Exec(
		departmentFormInput.GFKLocation,
		departmentFormInput.DepartmentCode,
		departmentFormInput.DepartmentName,
	)

	log.Println(res)

	return err
}

func CreateNewLocation(locationFormInput *costcenter.Location, dbBaa *sql.DB) error {

	insertNewLocationStr := `
	INSERT INTO baa_application.operation.location (
		fk_division
		,location_code
		,name
		) VALUES (@p1,@p2,@p3)`
	insertNewLocation, err := dbBaa.Prepare(insertNewLocationStr)

	res, err := insertNewLocation.Exec(
		locationFormInput.FKDivision,
		locationFormInput.LocationCode,
		locationFormInput.LocationName,
	)

	log.Println(res)

	return err
}

// CHANGE TO ID COST CENTER: join with cost_center table to get name

// GetPendingPurchaseRequest fetches pending purchase requests from baa_application.operation.purchase_request
func GetPendingPurchaseRequest(dbBaa *sql.DB) []*purchaserequestforminput.PurchaseRequestFormInput {
	rows, err := dbBaa.Query(` SELECT 
		pr.id_purchase_request
		,f.name cost_center  
		,pr.initiator  
		,pr.pr_type  
		,pr.cost_category  
		,pr.invoice_number  
		,CONVERT(VARCHAR(50), pr.invoice_date, 101) invoice_date
		,pr.vendor_name  
		,pr.item_description  
    	,CAST(ROUND(pr.unit_price ,2) as numeric(36,2)) unit_price
		,CAST(ROUND(pr.vat_unit_price ,2) as numeric(36,2)) vat_unit_price
		,pr.quantity
		,pr.payment_term 
		,pr.payment_installment 
		,pr.payment_center  
		,pr.payment_type 
    	,CAST(ROUND(pr.invoice_total ,2) as numeric(36,2)) invoice_total
		,CAST(ROUND(pr.vat_invoice_total ,2) as numeric(36,2)) vat_invoice_total
		FROM baa_application.operation.purchase_request pr
    JOIN baa_application.operation.func f
    ON pr.gfk_cost_center = f.gid_function
		WHERE pr.purchase_request_status = 'pending'`)

	// We return incase of an error, and defer the closing of the row structure
	checkError(err)
	//defer rows.Close()

	// Create the data structure that is returned from the function.
	purchaseRequestFormInputTable := []*purchaserequestforminput.PurchaseRequestFormInput{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a purchaseRequestFormInput,
		purchaseRequestFormInput := &purchaserequestforminput.PurchaseRequestFormInput{}
		// Populate the attributes of the purchaseRequestFormInput,
		// and return incase of an error
		err := rows.Scan(
			&purchaseRequestFormInput.IDPurchaseRequest,
			&purchaseRequestFormInput.CostCenter,
			&purchaseRequestFormInput.Initiator,
			&purchaseRequestFormInput.PrType,
			&purchaseRequestFormInput.CostCategory,
			&purchaseRequestFormInput.InvoiceNumber,
			&purchaseRequestFormInput.InvoiceDate,
			&purchaseRequestFormInput.VendorName,
			&purchaseRequestFormInput.ItemDescription,
			&purchaseRequestFormInput.UnitPrice,
			&purchaseRequestFormInput.VatUnitPrice,
			&purchaseRequestFormInput.Quantity,
			&purchaseRequestFormInput.PaymentTerm,
			&purchaseRequestFormInput.PaymentInstallment,
			&purchaseRequestFormInput.PaymentCenter,
			&purchaseRequestFormInput.PaymentType,
			&purchaseRequestFormInput.InvoiceTotal,
			&purchaseRequestFormInput.VatInvoiceTotal)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		purchaseRequestFormInputTable = append(purchaseRequestFormInputTable, purchaseRequestFormInput)
	}

	return purchaseRequestFormInputTable

}

func GetAvailableCostCenter(dbBaa *sql.DB, iDUser string) []*costcenter.CostCenter {

	stmt, err := dbBaa.Prepare(`
	SELECT 
	ccv.id_cost_center
	,ccv.cost_center_name
		
	 FROM baa_application.operation.cost_center_view ccv

	 WHERE ccv.gfk_department IN (
		SELECT pda.gfk_department 
	 	FROM baa_application.operation.pr_department_access pda
	 	WHERE pda.fk_user =@p1)

  	OR ccv.gfk_location IN (
		SELECT pla.gfk_location 
	 	FROM baa_application.operation.pr_location_access pla
	 	WHERE pla.fk_user =@p1)

  	OR ccv.fk_division IN (
		SELECT pdi.fk_division 
	 	FROM baa_application.operation.pr_division_access pdi
	 	WHERE pdi.fk_user =@p1)
	 `)
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query(iDUser)
	checkError(err)
	defer rows.Close()

	costCenterTable := []*costcenter.CostCenter{}

	for rows.Next() {
		// For each row returned by the table, create a pointer to a CostCenter,
		costCenter := &costcenter.CostCenter{}
		// Populate the attributes of the CostCenter,
		// and return incase of an error
		err := rows.Scan(
			&costCenter.GIDFunction,
			&costCenter.FunctionName,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		costCenterTable = append(costCenterTable, costCenter)
	}

	return costCenterTable

}

// GetIDPurchaseRequest fetches iDPurchaseRequestTable from baa_application.operation.purchase_request representing all pending requests
func GetIDPurchaseRequest(dbBaa *sql.DB) []string {

	// store iDPurchaseRequestQuery in a string
	iDPurchaseRequestQuery := `SELECT 
	pr.id_purchase_request 
	FROM baa_application.operation.purchase_request pr
	WHERE pr.purchase_request_status = 'pending'`

	// write iDPurchaseRequestQuery result to an array of fields.InputChoice , this array of rows represents iDPurchaseRequestTable
	var iDPurchaseRequest string
	var iDPurchaseRequestTable []string
	rows, err := dbBaa.Query(iDPurchaseRequestQuery)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&iDPurchaseRequest)
		checkError(err)
		iDPurchaseRequestTable = append(iDPurchaseRequestTable, iDPurchaseRequest)
	}
	return iDPurchaseRequestTable
}

func ConvertPurchaseRequestToPurchaseOrder(iDPurchaseRequest string, dbBaa *sql.DB) error {

	convertPurchaseRequestToPurchaseOrderStr := `
	UPDATE baa_application.operation.purchase_request
	SET baa_application.operation.purchase_request.purchase_request_status = 'approved'
	WHERE baa_application.operation.purchase_request.id_purchase_request = @p1`

	convertPurchaseRequestToPurchaseOrder, err := dbBaa.Prepare(convertPurchaseRequestToPurchaseOrderStr)

	res, err := convertPurchaseRequestToPurchaseOrder.Exec(iDPurchaseRequest)

	log.Println(res)

	return err
}

func ConvertPurchaseRequestToRejectedPurchaseRequest(iDPurchaseRequest string, dbBaa *sql.DB) error {

	convertPurchaseRequestToRejectedPurchaseRequestStr := `
	UPDATE baa_application.operation.purchase_request
	SET baa_application.operation.purchase_request.purchase_request_status = 'rejected'
	WHERE baa_application.operation.purchase_request.id_purchase_request = @p1`

	convertPurchaseRequestToRejectedPurchaseRequest, err := dbBaa.Prepare(convertPurchaseRequestToRejectedPurchaseRequestStr)

	res, err := convertPurchaseRequestToRejectedPurchaseRequest.Exec(iDPurchaseRequest)

	log.Println(res)

	return err
}

func GetUserInfo(user *useraccess.User, dbBaa *sql.DB) {

	// store userQuery in a string
	userQuery := `SELECT 
	pu.id_user
	,pu.name
	,pu.access
	FROM baa_application.operation.pr_user pu
	WHERE pu.email = @p1`

	err := dbBaa.QueryRow(userQuery, user.Email).Scan(&user.IDUser, &user.Name, &user.Access)
	if err != nil {
		if err.Error() != `sql: no rows in result set` {
			log.Fatal(err.Error())
		}

	}

}

func GetListOfFunctionName(dbBaa *sql.DB) string {

	// store query in a string
	query := `SELECT fu.name FROM baa_application.operation.func fu`

	var functionName, listOfFunctionName string
	rows, err := dbBaa.Query(query)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&functionName)
		checkError(err)
		listOfFunctionName += functionName + `,`
	}
	listOfFunctionName = listOfFunctionName[:len(listOfFunctionName)-1]
	return listOfFunctionName

}

func GetNextFunctionCode(dbBaa *sql.DB) string {

	// store query in a string
	query := `
	SELECT 
	CASE 
	WHEN COALESCE(MAX(fu.function_code)+1,0) < 10  THEN CONCAT('0',CAST(COALESCE(MAX(fu.function_code),0)+1 AS VARCHAR))
	ELSE  CAST(MAX(fu.function_code+1) AS VARCHAR) 
	END  next_function_code
  	FROM baa_application.operation.func fu`

	var nextFunctionCode string
	err := dbBaa.QueryRow(query).Scan(&nextFunctionCode)
	checkError(err)
	return nextFunctionCode

}

func GetListOfDepartmentName(dbBaa *sql.DB) string {

	// store query in a string
	query := `SELECT de.name FROM baa_application.operation.department de`

	var departmentName, listOfDepartmentName string
	rows, err := dbBaa.Query(query)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&departmentName)
		checkError(err)
		listOfDepartmentName += departmentName + `,`
	}
	listOfDepartmentName = listOfDepartmentName[:len(listOfDepartmentName)-1]
	return listOfDepartmentName

}

func GetNextDepartmentCode(dbBaa *sql.DB) string {

	// store query in a string
	query := `
	SELECT 
  	CASE 
  	WHEN COALESCE(MAX(d.department_code)+1,0) < 10  THEN CONCAT('0',CAST(COALESCE(MAX(d.department_code),0)+1 AS VARCHAR))
  	ELSE  CAST(MAX(d.department_code+1) AS VARCHAR) 
  	END  next_department_code
	FROM baa_application.operation.department d`

	var nextDepartmentCode string
	err := dbBaa.QueryRow(query).Scan(&nextDepartmentCode)
	checkError(err)
	return nextDepartmentCode

}

func GetListOfLocationName(dbBaa *sql.DB) string {

	// store query in a string
	query := `SELECT lo.name FROM baa_application.operation.location lo`

	var locationName, listOfLocationName string
	rows, err := dbBaa.Query(query)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&locationName)
		checkError(err)
		listOfLocationName += locationName + `,`
	}
	listOfLocationName = listOfLocationName[:len(listOfLocationName)-1]
	return listOfLocationName

}

func GetNextLocationCode(dbBaa *sql.DB) string {

	// store query in a string
	query := `
	SELECT 
	CASE 
	WHEN COALESCE(MAX(l.location_code)+1,0) < 10 THEN CONCAT('0',CAST(COALESCE(MAX(l.location_code),0)+1 AS VARCHAR))
	ELSE  CAST(MAX(l.location_code+1) AS VARCHAR) 
	END  next_location_code
  	FROM baa_application.operation.location l`

	var nextLocationCode string
	err := dbBaa.QueryRow(query).Scan(&nextLocationCode)
	checkError(err)
	return nextLocationCode

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
