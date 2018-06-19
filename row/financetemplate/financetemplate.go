package financetemplate

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type FinanceTemplate struct {
	InvoiceTotalWithVAT string `csv:"بستانكار(ريال)"`
	InvoiceCount        string `csv:"تعداد قبض"`
	ItemDescription     string `csv:"شرح فاکتور"`
	CostCategoryCode    string `csv:"کد  معین"`
	CostCategoryName    string `csv:"شرح معین"`
	GFKCostCenter       string `csv:"مرکز هزینه/د‍‍‍پارتمان"`
	InvoiceNumber       string `csv:"شماره فاكتور"`
	InvoiceDate         string `csv:"تاريخ فاكتور"`
	NationalID          string `csv:"كدملي/كد اقتصادي"`
	VendorCode          string `csv:"کد شناور"`
	VendorName          string `csv:"نام فروشنده"`
	RowNumber           string `csv:"ردیف"`

	StartDate string `jon:"start_date"`
	EndDate   string `jon:"end_date"`

	Error string
}

// Validate validates the data of the purchase request sent by the user
func (financeTemplate *FinanceTemplate) Validate() bool {

	financeTemplate.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(financeTemplate,
		validation.Field(&financeTemplate.StartDate, validation.Required, validation.Date("1/2/2006")),
		validation.Field(&financeTemplate.EndDate, validation.Required, validation.Date("1/2/2006")),
	)

	// add potential error text to financeTemplate.Error
	if err != nil {
		financeTemplate.Error = err.Error()
	}

	// return true if no error, false otherwise
	return financeTemplate.Error == ""
}

// Render the web page itself given the html template and the financeTemplate
func (financeTemplate *FinanceTemplate) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the financeTemplate
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`StartDate`: financeTemplate.StartDate,
		`EndDate`:   financeTemplate.EndDate,

		`Error`: financeTemplate.Error,
	})
	handleErr(c, err)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
