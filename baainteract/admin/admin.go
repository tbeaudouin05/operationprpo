package admin

import (
	"database/sql"
	"log"

	"github.com/thomas-bamilo/operation/operationprpo/row/purchaserequestforminput"
)

// GetPendingPurchaseRequest fetches pending purchase requests from baa_application.operation.purchase_request
func GetPendingPurchaseRequest(dbBaa *sql.DB) []*purchaserequestforminput.PurchaseRequestFormInput {
	rows, err := dbBaa.Query(`
		SELECT 
		pr.id_purchase_request
		,f.name cost_center  
		,pr.initiator  
		,pr.pr_type  
		,pcc.name_fa cost_category  
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

		JOIN baa_application.operation.pr_cost_category pcc
		ON pcc.id_cost_category = pr.fk_cost_category

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

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
