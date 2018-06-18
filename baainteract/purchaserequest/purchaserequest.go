package purchaserequest

import (
	"database/sql"
	"log"

	"github.com/thomas-bamilo/operation/operationprpo/row/costcenter"

	"github.com/thomas-bamilo/operation/operationprpo/row/purchaserequestforminput"
	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"
)

// LoadPurchaseRequestToDb loads purchase request to baa_application.operation.purchase_request
func LoadPurchaseRequestToDb(purchaseRequestFormInput *purchaserequestforminput.PurchaseRequestFormInput, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertPurchaseRequestStr := `
INSERT INTO baa_application.operation.purchase_request (
	gfk_cost_center  
	,initiator  
	,pr_type
	,cost_type  
	,fk_cost_category  
	,invoice_number  
	,invoice_date  
	,fk_vendor  
	,item_description  
	,unit_price  
	,vat_unit_price  
	,quantity  
	,payment_term  
	,payment_installment
	,payment_center   
	,payment_type
	,purchase_request_status) 
VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13,@p14,@p15,@p16,'pending')`
	insertPurchaseRequest, err := dbBaa.Prepare(insertPurchaseRequestStr)

	res, err := insertPurchaseRequest.Exec(
		purchaseRequestFormInput.CostCenter,
		purchaseRequestFormInput.Initiator,
		purchaseRequestFormInput.PrType,
		purchaseRequestFormInput.CostType,
		purchaseRequestFormInput.CostCategory,
		purchaseRequestFormInput.InvoiceNumber,
		purchaseRequestFormInput.InvoiceDate,
		purchaseRequestFormInput.FKVendor,
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

func GetAvailableCostCenter(dbBaa *sql.DB, iDUser string) []*costcenter.CostCenter {

	stmt, err := dbBaa.Prepare(`
	SELECT 
	ccv.gid_cost_center
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

func GetCostCategory(dbBaa *sql.DB) []*purchaserequestforminput.CostCategory {

	stmt, err := dbBaa.Prepare(`
		SELECT 

		pcc.id_cost_category
		,pcc.name_fa
		  
		FROM baa_application.operation.pr_cost_category pcc
	 `)
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkError(err)
	defer rows.Close()

	costCategoryTable := []*purchaserequestforminput.CostCategory{}

	for rows.Next() {
		// For each row returned by the table, create a pointer to a CostCategory,
		costCategory := &purchaserequestforminput.CostCategory{}
		// Populate the attributes of the CostCategory,
		// and return incase of an error
		err := rows.Scan(
			&costCategory.IDCostCategory,
			&costCategory.NameFa,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		costCategoryTable = append(costCategoryTable, costCategory)
	}

	return costCategoryTable

}

func GetVendor(dbBaa *sql.DB) []*purchaserequestforminput.Vendor {

	stmt, err := dbBaa.Prepare(`
		SELECT 
  
		v.id_vendor
		,v.name
		
	  FROM baa_application.finance.vendor v
	 `)
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkError(err)
	defer rows.Close()

	vendorTable := []*purchaserequestforminput.Vendor{}

	for rows.Next() {
		// For each row returned by the table, create a pointer to a Vendor,
		vendor := &purchaserequestforminput.Vendor{}
		// Populate the attributes of the Vendor,
		// and return incase of an error
		err := rows.Scan(
			&vendor.IDVendor,
			&vendor.Name,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		vendorTable = append(vendorTable, vendor)
	}

	return vendorTable

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

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
