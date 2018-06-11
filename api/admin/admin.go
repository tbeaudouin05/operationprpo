package admin

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/operation/operationprpo/api/oauth/authorize"
	ad_baainteract "github.com/thomas-bamilo/operation/operationprpo/baainteract/admin"
	"github.com/thomas-bamilo/operation/operationprpo/row/adminchoice"
	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"
	"github.com/thomas-bamilo/sql/connectdb"
)

var user useraccess.User

// Start loads the admin web page - GET request
func Start(c *gin.Context) {

	authorize.AuthorizePrAdmin(c, &user)

	// empty adminChoice to start
	adminChoice := &adminchoice.AdminChoice{}

	// render the web page itself given the html template and the adminChoice
	adminChoice.Render(c, `template/admin/admin.html`)
}

// StartIDPurchaseRequest populates the admin form with iDPurchaseRequest options - GET request
func StartIDPurchaseRequest(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	// GetIDPurchaseRequest queries baa_application.operation.purchase_request to return all pending iDPurchaseRequest
	iDPurchaseRequestTable := ad_baainteract.GetIDPurchaseRequest(dbBaa)

	//Convert the `iDPurchaseRequestTable` variable to json
	iDPurchaseRequestTableByte, err := json.Marshal(iDPurchaseRequestTable)

	// If there is an error, print it to the console, and return a server error response to the user
	handleErr(c, err)

	// If all goes well, write the JSON list of iDPurchaseRequestTable to the response
	c.Writer.Write(iDPurchaseRequestTableByte)
}

// StartPendingPurchaseRequest populates the admin web page with pending purchase requests - GET request
func StartPendingPurchaseRequest(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	// GetPendingPurchaseRequest queries baa_application.operation.purchase_request to return all pending purchase requests
	purchaseRequestFormInputTable := ad_baainteract.GetPendingPurchaseRequest(dbBaa)

	//Convert the `purchaseRequestFormInputTable` variable to json
	purchaseRequestFormInputTableByte, err := json.Marshal(purchaseRequestFormInputTable)
	handleErr(c, err)

	// If all goes well, write the JSON list of purchaseRequestFormInputTable to the response
	c.Writer.Write(purchaseRequestFormInputTableByte)
}

// AcceptRejectPurchaseRequest records admin user inputs to accept or reject a purchase request
// and accept or reject the given purchase request - POST request
func AcceptRejectPurchaseRequest(c *gin.Context) {

	authorize.AuthorizePrAdmin(c, &user)

	// pass all the form values input by the user and the form as adminChoice
	// since we want to validate these values and upload them to database
	// in case validation fails, we also want to return these values to the form for good user experience
	adminChoice := &adminchoice.AdminChoice{
		IDPurchaseRequest: c.Request.FormValue(`iDPurchaseRequest`),
		AcceptReject:      c.Request.FormValue(`acceptReject`),
	}

	// Validate validates the adminChoice form user inputs
	// if validation fails, reload the purchase request form page with the initial user inputs and error messages
	if adminChoice.Validate() == false {
		adminChoice.Render(c, `template/admin/admin.html`)
		return
	}

	if adminChoice.AcceptReject == `Accept` {
		dbBaa := connectdb.ConnectToBaa()
		err := ad_baainteract.ConvertPurchaseRequestToPurchaseOrder(adminChoice.IDPurchaseRequest, dbBaa)
		handleErr(c, err)
	}

	if adminChoice.AcceptReject == `Reject` {
		dbBaa := connectdb.ConnectToBaa()
		err := ad_baainteract.ConvertPurchaseRequestToRejectedPurchaseRequest(adminChoice.IDPurchaseRequest, dbBaa)
		handleErr(c, err)
	}

	// if everything goes well, reload web page
	http.Redirect(c.Writer, c.Request, `/admin`, http.StatusSeeOther)

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
