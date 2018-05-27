package purchaserequestforminput

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// PurchaseRequestFormInput represents one purchase request
// it also includes potential error or form to include into the web page used to collect purchase requests
type PurchaseRequestFormInput struct {
	IDPurchaseRequest string `json:"id_purchase_request"`
	CostCenter        string `json:"cost_center"`
	Initiator         string `json:"initiator"`
	PrType            string `json:"pr_type"`
	CostCategory      string `json:"cost_category"`
	InvoiceNumber     string `json:"invoice_number"`
	InvoiceDate       string `json:"invoice_date"`
	VendorName        string `json:"vendor_name"`
	ItemDescription   string `json:"item_description"`
	UnitPrice         string `json:"unit_price"`
	VatUnitPrice      string `json:"vat_unit_price"`
	Quantity          string `json:"quantity"`
	PaymentTerm       string `json:"payment_term"`
	PaymentCenter     string `json:"payment_center"`
	AlreadyPaid       string `json:"already_paid"`
	PaymentType       string `json:"payment_type"`
	InvoiceTotal      string `json:"invoice_total"`
	VatInvoiceTotal   string `json:"vat_invoice_total"`

	Error string
}

// Validate validates the data of the purchase request sent by the user
func (purchaseRequestFormInput *PurchaseRequestFormInput) Validate() bool {

	purchaseRequestFormInput.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(purchaseRequestFormInput,
		validation.Field(&purchaseRequestFormInput.CostCenter, validation.Required),
		validation.Field(&purchaseRequestFormInput.Initiator, validation.Required),
		validation.Field(&purchaseRequestFormInput.PrType, validation.Required),
		validation.Field(&purchaseRequestFormInput.CostCategory, validation.Required),
		validation.Field(&purchaseRequestFormInput.InvoiceNumber, validation.Required),
		validation.Field(&purchaseRequestFormInput.InvoiceDate, validation.Required, validation.Date("1/2/2006")),
		validation.Field(&purchaseRequestFormInput.VendorName, validation.Required),
		validation.Field(&purchaseRequestFormInput.ItemDescription, validation.Required),
		validation.Field(&purchaseRequestFormInput.UnitPrice, validation.Required, is.Float),
		validation.Field(&purchaseRequestFormInput.VatUnitPrice, validation.Required, is.Float),
		validation.Field(&purchaseRequestFormInput.Quantity, validation.Required, is.Int),
		validation.Field(&purchaseRequestFormInput.PaymentTerm, validation.Required),
		validation.Field(&purchaseRequestFormInput.PaymentCenter, validation.Required),
		validation.Field(&purchaseRequestFormInput.AlreadyPaid, validation.Required),
		validation.Field(&purchaseRequestFormInput.PaymentType, validation.Required),
	)

	// add potential error text to purchaseRequestFormInput.Error
	if err != nil {
		purchaseRequestFormInput.Error = err.Error()
	}

	// return true if no error, false otherwise
	return purchaseRequestFormInput.Error == ""
}

// Render the web page itself given the html template and the purchaseRequestFormInput
func (purchaseRequestFormInput *PurchaseRequestFormInput) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the purchaseRequestFormInput
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`IDPurchaseRequest`: purchaseRequestFormInput.IDPurchaseRequest,
		`CostCenter`:        purchaseRequestFormInput.CostCenter,
		`Initiator`:         purchaseRequestFormInput.Initiator,
		`PrType`:            purchaseRequestFormInput.PrType,
		`CostCategory`:      purchaseRequestFormInput.CostCategory,
		`InvoiceNumber`:     purchaseRequestFormInput.InvoiceNumber,
		`InvoiceDate`:       purchaseRequestFormInput.InvoiceDate,
		`VendorName`:        purchaseRequestFormInput.VendorName,
		`ItemDescription`:   purchaseRequestFormInput.ItemDescription,
		`UnitPrice`:         purchaseRequestFormInput.UnitPrice,
		`VatUnitPrice`:      purchaseRequestFormInput.VatUnitPrice,
		`Quantity`:          purchaseRequestFormInput.Quantity,
		`PaymentTerm`:       purchaseRequestFormInput.PaymentTerm,
		`PaymentCenter`:     purchaseRequestFormInput.PaymentCenter,
		`AlreadyPaid`:       purchaseRequestFormInput.AlreadyPaid,
		`PaymentType`:       purchaseRequestFormInput.PaymentType,
		`InvoiceTotal`:      purchaseRequestFormInput.InvoiceTotal,
		`VatInvoiceTotal`:   purchaseRequestFormInput.VatInvoiceTotal,
		`Error`:             purchaseRequestFormInput.Error,
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
