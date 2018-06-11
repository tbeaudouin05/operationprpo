package purchaserequestform

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/operation/operationprpo/api/oauth/authorize"
	pr_baainteract "github.com/thomas-bamilo/operation/operationprpo/baainteract/purchaserequest"
	"github.com/thomas-bamilo/operation/operationprpo/row/purchaserequestforminput"
	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"
	"github.com/thomas-bamilo/sql/connectdb"
)

var user useraccess.User

// Start loads the purchase request form web page - GET request
func Start(c *gin.Context) {

	authorize.Authorize(c, &user)

	// only pass form as purchaseRequestFormInput since we only want a blank form at start
	purchaseRequestFormInput := &purchaserequestforminput.PurchaseRequestFormInput{}

	// render the web page itself given the html template and the purchaseRequestFormInput
	purchaseRequestFormInput.Render(c, `template/purchaserequest/purchaserequest.html`)
}

// StartInvoiceDate populates the purchase request form with invoice dates options - GET request
func StartInvoiceDate(c *gin.Context) {

	// GetPendingPurchaseRequest queries baa_application.operation.purchase_request to return all pending purchase requests
	var invoiceDateTable []string

	today := time.Now()
	for i := 1; i <= 60; i++ {
		invoiceDateTable = append(invoiceDateTable, today.AddDate(0, 0, i-30).Format(`1/2/2006`))
	}

	//Convert the `invoiceDateTable` variable to json
	invoiceDateTableByte, err := json.Marshal(invoiceDateTable)
	handleErr(c, err)

	// If all goes well, write the JSON list of invoiceDateTable to the response
	c.Writer.Write(invoiceDateTableByte)
}

// StartCostCategory populates the admin web page with pending purchase requests - GET request
func StartCostCategory(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	costCenterTable := pr_baainteract.GetCostCategory(dbBaa)

	//Convert the `costCenterTable` variable to json
	costCenterTableByte, err := json.Marshal(costCenterTable)
	handleErr(c, err)

	// If all goes well, write the JSON list of costCenterTable to the response
	c.Writer.Write(costCenterTableByte)
}

// StartAvailableCostCenter populates the admin web page with pending purchase requests - GET request
func StartAvailableCostCenter(c *gin.Context) {

	// DO NOT PUT AUTHORIZE HERE OTHERWISE CHANGES SESSION AND BREAKS EVERYTHING

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	costCenterTable := pr_baainteract.GetAvailableCostCenter(dbBaa, user.IDUser)

	//Convert the `costCenterTable` variable to json
	costCenterTableByte, err := json.Marshal(costCenterTable)
	handleErr(c, err)

	// If all goes well, write the JSON list of costCenterTable to the response
	c.Writer.Write(costCenterTableByte)
}

// AnswerForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerForm(c *gin.Context) {

	authorize.Authorize(c, &user)

	r := c.Request

	// pass all the form values input by the user
	// since we want to validate these values and upload them to database
	// in case validation fails, we also want to return these values to the form for good user experience
	purchaseRequestFormInput := &purchaserequestforminput.PurchaseRequestFormInput{
		CostCenter:         r.FormValue(`costCenter`),
		Initiator:          user.Name,
		PrType:             r.FormValue(`prType`),
		CostCategory:       r.FormValue(`costCategory`),
		InvoiceNumber:      r.FormValue(`invoiceNumber`),
		InvoiceDate:        r.FormValue(`invoiceDate`),
		VendorName:         r.FormValue(`vendorName`),
		ItemDescription:    r.FormValue(`itemDescription`),
		UnitPrice:          r.FormValue(`unitPrice`),
		VatUnitPrice:       r.FormValue(`vatUnitPrice`),
		Quantity:           r.FormValue(`quantity`),
		PaymentTerm:        r.FormValue(`paymentTerm`),
		PaymentInstallment: r.FormValue(`paymentInstallment`),
		PaymentCenter:      r.FormValue(`paymentCenter`),
		PaymentType:        r.FormValue(`paymentType`),

		IsAnotherItem: r.FormValue(`isAnotherItem`),
	}

	// Validate validates the purchaseRequestFormInput form user inputs
	// if validation fails, reload the purchase request form page with the initial user inputs and error messages
	if purchaseRequestFormInput.Validate() == false {
		purchaseRequestFormInput.Render(c, `template/purchaserequest/purchaserequest.html`)
		return
	}

	// LoadToDb uploads the purchase request form user inputs (= purchaseRequestFormInput) to database
	dbBaa := connectdb.ConnectToBaa()
	err := pr_baainteract.LoadPurchaseRequestToDb(purchaseRequestFormInput, dbBaa)
	handleErr(c, err)

	// if isAnotherItem = yes ie the user wants to make a similar purchase request
	// then reload the page with same information + a success message
	if purchaseRequestFormInput.IsAnotherItem == `yes` {
		purchaseRequestFormInput.Success = `Purchase request successful: please add another item`
		purchaseRequestFormInput.Render(c, `template/purchaserequest/purchaserequest.html`)
		return
	}

	// if everything goes well, redirect user to confirmation web page
	http.Redirect(c.Writer, r, `/purchaserequest/purchaserequestconfirmation`, http.StatusSeeOther)
}

// ConfirmForm loads the purchase request confirmation web page - GET request
func ConfirmForm(c *gin.Context) {

	// render confirmation web page
	render(c, `template/purchaserequest/purchaserequestconfirmation.html`)
}

// Render the web page itself given the html template - no parameter
func render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate without parameter
	err = tmpl.Execute(c.Writer, nil)
	handleErr(c, err)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
