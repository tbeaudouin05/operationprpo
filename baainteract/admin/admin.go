package admin

import (
	"database/sql"
	"log"

	"github.com/thomas-bamilo/operation/operationprpo/row/financetemplate"
	"github.com/thomas-bamilo/operation/operationprpo/row/purchaserequestforminput"
)

// GetPendingPurchaseRequest fetches pending purchase requests from baa_application.operation.purchase_request
func GetPendingPurchaseRequest(dbBaa *sql.DB) []*purchaserequestforminput.PurchaseRequestFormInput {
	rows, err := dbBaa.Query(`
		SELECT 
		pr.id_purchase_request
		,CONVERT(VARCHAR(50), pr.pr_timestamp, 100) pr_timestamp
		,f.name cost_center  
		,pr.initiator  
		,pr.pr_type
		,pr.cost_type  
		,pcc.name_fa cost_category  
		,pr.invoice_count
		,pr.invoice_number  
		,CONVERT(VARCHAR(50), pr.invoice_date, 101) invoice_date
		,v.name vendor_name 
		,pr.item_description  
		,replace(convert(varchar,convert(Money,  pr.unit_price),1),'.00','') unit_price
		,replace(convert(varchar,convert(Money,  pr.vat_unit_price),1),'.00','') vat_unit_price
		,pr.quantity
		,pr.payment_term 
		,pr.payment_installment 
		,pr.payment_center  
		,pr.payment_type 
		,replace(convert(varchar,convert(Money,  pr.invoice_total),1),'.00','') invoice_total
		,replace(convert(varchar,convert(Money,  pr.vat_invoice_total),1),'.00','') vat_invoice_total

		FROM baa_application.operation.purchase_request pr

    	JOIN baa_application.operation.func f
		ON pr.gfk_cost_center = f.gid_function

		JOIN baa_application.operation.pr_cost_category pcc
		ON pcc.id_cost_category = pr.fk_cost_category

		JOIN baa_application.finance.vendor v
    	ON pr.fk_vendor = v.id_vendor

		WHERE pr.purchase_request_status = 'pending'`)

	// We return incase of an error, and defer the closing of the row structure
	checkError(err)
	defer rows.Close()

	// Create the data structure that is returned from the function.
	purchaseRequestFormInputTable := []*purchaserequestforminput.PurchaseRequestFormInput{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a purchaseRequestFormInput,
		purchaseRequestFormInput := &purchaserequestforminput.PurchaseRequestFormInput{}
		// Populate the attributes of the purchaseRequestFormInput,
		// and return incase of an error
		err := rows.Scan(
			&purchaseRequestFormInput.IDPurchaseRequest,
			&purchaseRequestFormInput.Timestamp,
			&purchaseRequestFormInput.CostCenter,
			&purchaseRequestFormInput.Initiator,
			&purchaseRequestFormInput.PrType,
			&purchaseRequestFormInput.CostType,
			&purchaseRequestFormInput.CostCategory,
			&purchaseRequestFormInput.NumberOfInvoice,
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

// GetApprovedPurchaseRequest fetches approved purchase requests from baa_application.operation.purchase_request
func GetApprovedPurchaseRequest(dbBaa *sql.DB) []*purchaserequestforminput.PurchaseRequestFormInput {
	rows, err := dbBaa.Query(`
		SELECT 
		pr.id_purchase_request
		,CONVERT(VARCHAR(50), pr.pr_timestamp, 100) pr_timestamp
		,f.name cost_center  
		,pr.initiator  
		,pr.pr_type
		,pr.cost_type  
		,pcc.name_fa cost_category  
		,pr.invoice_count
		,pr.invoice_number  
		,CONVERT(VARCHAR(50), pr.invoice_date, 101) invoice_date
		,v.name vendor_name 
		,pr.item_description  
		,replace(convert(varchar,convert(Money,  pr.unit_price),1),'.00','') unit_price
		,replace(convert(varchar,convert(Money,  pr.vat_unit_price),1),'.00','') vat_unit_price
		,pr.quantity
		,pr.payment_term 
		,pr.payment_installment 
		,pr.payment_center  
		,pr.payment_type 
		,replace(convert(varchar,convert(Money,  pr.invoice_total),1),'.00','') invoice_total
		,replace(convert(varchar,convert(Money,  pr.vat_invoice_total),1),'.00','') vat_invoice_total

		FROM baa_application.operation.purchase_request pr

    	JOIN baa_application.operation.func f
		ON pr.gfk_cost_center = f.gid_function

		JOIN baa_application.operation.pr_cost_category pcc
		ON pcc.id_cost_category = pr.fk_cost_category

		JOIN baa_application.finance.vendor v
    	ON pr.fk_vendor = v.id_vendor

		WHERE pr.purchase_request_status = 'approved'`)

	// We return incase of an error, and defer the closing of the row structure
	checkError(err)
	defer rows.Close()

	// Create the data structure that is returned from the function.
	purchaseRequestFormInputTable := []*purchaserequestforminput.PurchaseRequestFormInput{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a purchaseRequestFormInput,
		purchaseRequestFormInput := &purchaserequestforminput.PurchaseRequestFormInput{}
		// Populate the attributes of the purchaseRequestFormInput,
		// and return incase of an error
		err := rows.Scan(
			&purchaseRequestFormInput.IDPurchaseRequest,
			&purchaseRequestFormInput.Timestamp,
			&purchaseRequestFormInput.CostCenter,
			&purchaseRequestFormInput.Initiator,
			&purchaseRequestFormInput.PrType,
			&purchaseRequestFormInput.CostType,
			&purchaseRequestFormInput.CostCategory,
			&purchaseRequestFormInput.NumberOfInvoice,
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

func DownloadFinanceTemplate(dbBaa *sql.DB, financeTemplateForm financetemplate.FinanceTemplate) []*financetemplate.FinanceTemplate {

	query := `
	SELECT 

	replace(convert(varchar,convert(Money,  pr.invoice_total + vat_invoice_total),1),'.00','')  'بستانكار(ريال)'
	,pr.invoice_count 'تعداد قبض'
	,pr.item_description 'شرح فاکتور'
	,pcc.code 'کد  معین'
	,pcc.name_fa 'شرح معین'
	,pr.gfk_cost_center 'مرکز هزینه/د‍‍‍پارتمان'
	,pr.invoice_number 'شماره فاكتور'
	,CONVERT(VARCHAR(50), pr.invoice_date, 101) 'تاريخ فاكتور'
	,'' 'كدملي/كد اقتصادي'
	,v.code 'کد شناور'
	,v.name 'نام فروشنده'
	,ROW_NUMBER() OVER (ORDER BY(pr.invoice_date))  'ردیف'
	
	FROM baa_application.operation.purchase_request pr
	
	JOIN baa_application.operation.pr_cost_category pcc
	ON pr.fk_cost_category = pcc.id_cost_category
	
	JOIN baa_application.operation.cost_center_view ccv
	ON ccv.gid_cost_center = pr.gfk_cost_center
	
	JOIN baa_application.finance.vendor v
	ON v.id_vendor = pr.fk_vendor
	
	WHERE pr.purchase_request_status = 'approved'
	AND pr.pr_timestamp BETWEEN @p1 AND @p2
	`

	rows, err := dbBaa.Query(query, financeTemplateForm.StartDate, financeTemplateForm.EndDate)

	// We return incase of an error, and defer the closing of the row structure
	checkError(err)
	defer rows.Close()

	// Create the data structure that is returned from the function.
	financeTemplateTable := []*financetemplate.FinanceTemplate{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a financeTemplate,
		financeTemplate := &financetemplate.FinanceTemplate{}
		// Populate the attributes of the financeTemplate,
		// and return incase of an error
		err := rows.Scan(
			&financeTemplate.InvoiceTotalWithVAT,
			&financeTemplate.InvoiceCount,
			&financeTemplate.ItemDescription,
			&financeTemplate.CostCategoryCode,
			&financeTemplate.CostCategoryName,
			&financeTemplate.GFKCostCenter,
			&financeTemplate.InvoiceNumber,
			&financeTemplate.InvoiceDate,
			&financeTemplate.NationalID,
			&financeTemplate.VendorCode,
			&financeTemplate.VendorName,
			&financeTemplate.RowNumber,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		financeTemplateTable = append(financeTemplateTable, financeTemplate)
	}

	return financeTemplateTable

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
