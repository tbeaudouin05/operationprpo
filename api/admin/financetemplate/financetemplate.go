package financetemplate

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/operation/operationprpo/api/oauth/authorize"
	ad_baainteract "github.com/thomas-bamilo/operation/operationprpo/baainteract/admin"
	"github.com/thomas-bamilo/operation/operationprpo/row/financetemplate"
	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"
	"github.com/thomas-bamilo/sql/connectdb"
)

var user useraccess.User
var financeTemplateForm financetemplate.FinanceTemplate

// Start loads the admin web page - GET request
func Start(c *gin.Context) {

	authorize.AuthorizePrAdmin(c, &user)

	// empty financeTemplate to start
	financeTemplate := &financetemplate.FinanceTemplate{}

	// render the web page itself given the html template and the financeTemplate
	financeTemplate.Render(c, `template/admin/financetemplate/financetemplate.html`)
}

// - GET request
func StartSuccess(c *gin.Context) {

	authorize.AuthorizePrAdmin(c, &user)

	// empty financeTemplate to start
	financeTemplate := &financetemplate.FinanceTemplate{}

	// render the web page itself given the html template and the financeTemplate
	financeTemplate.Render(c, `template/admin/financetemplate/financetemplatesuccess.html`)
}

// StartApprovedPurchaseRequest populates the admin web page with approved purchase requests - GET request
func StartApprovedPurchaseRequest(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	// GetApprovedPurchaseRequest queries baa_application.operation.purchase_request to return all approved purchase requests
	purchaseRequestFormInputTable := ad_baainteract.GetApprovedPurchaseRequest(dbBaa)

	//Convert the `purchaseRequestFormInputTable` variable to json
	purchaseRequestFormInputTableByte, err := json.Marshal(purchaseRequestFormInputTable)
	handleErr(c, err)

	// If all goes well, write the JSON list of purchaseRequestFormInputTable to the response
	c.Writer.Write(purchaseRequestFormInputTableByte)
}

// POST
func FinanceTemplateForm(c *gin.Context) {

	authorize.AuthorizePrAdmin(c, &user)

	r := c.Request

	startDate, err := time.Parse("2006-1-2", r.FormValue(`StartDate`))
	handleErr(c, err)
	startDateStr := startDate.Format("1/2/2006")
	endDate, err := time.Parse("2006-1-2", r.FormValue(`EndDate`))
	handleErr(c, err)
	endDateStr := endDate.Format("1/2/2006")

	financeTemplateForm = financetemplate.FinanceTemplate{
		StartDate: startDateStr,
		EndDate:   endDateStr,
	}

	if financeTemplateForm.Validate() == false {
		financeTemplateForm.Render(c, `template/admin/financetemplate/financetemplate.html`)
		return
	}

	// if everything goes well, load success page
	http.Redirect(c.Writer, c.Request, `/admin/financetemplatesuccess`, http.StatusSeeOther)

}

// GET
func DownloadFinanceTemplate(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	//get csv information
	financeTemplate := ad_baainteract.DownloadFinanceTemplate(dbBaa, financeTemplateForm)

	// download csv
	var financeTemplateStr [][]string

	financeTemplateStr = append(financeTemplateStr, []string{`بستانكار(ريال)`,
		`تعداد قبض`,
		`شرح فاکتور`,
		`کد  معین`,
		`شرح معین`,
		`مرکز هزینه/د‍‍‍پارتمان`,
		`شماره فاكتور`,
		`تاريخ فاكتور`,
		`كدملي/كد اقتصادي`,
		`کد شناور`,
		`نام فروشنده`,
		`ردیف`,
	})

	for i := 0; i < len(financeTemplate); i++ {

		financeTemplateStr = append(financeTemplateStr, []string{
			financeTemplate[i].InvoiceTotalWithVAT,
			financeTemplate[i].InvoiceCount,
			financeTemplate[i].ItemDescription,
			financeTemplate[i].CostCategoryCode,
			financeTemplate[i].CostCategoryName,
			financeTemplate[i].GFKCostCenter,
			financeTemplate[i].InvoiceNumber,
			financeTemplate[i].InvoiceDate,
			financeTemplate[i].NationalID,
			financeTemplate[i].VendorCode,
			financeTemplate[i].VendorName,
			financeTemplate[i].RowNumber,
		})
	}

	b := &bytes.Buffer{}   // creates IO Writer
	wr := csv.NewWriter(b) // creates a csv writer that uses the io buffer.

	for i := 0; i < len(financeTemplateStr); i++ {
		wr.Write(financeTemplateStr[i]) // converts array of string to comma seperated values for 1 row.
	}
	wr.Flush() // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=finance_template.csv")
	c.Data(http.StatusOK, "text/csv", b.Bytes())

}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
