package purchaserequestforminput

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// PurchaseRequestFormInput represents one purchase request
// it also includes potential error or form to include into the web page used to collect purchase requests
type PurchaseRequestFormInput struct {
	IDPurchaseRequest  string `json:"id_purchase_request"`
	Timestamp          string `json:"timestamp"`
	CostCenter         string `json:"cost_center"`
	Initiator          string `json:"initiator"`
	PrType             string `json:"pr_type"`
	CostType           string `json:"cost_type"`
	CostCategory       string `json:"cost_category"`
	InvoiceNumber      string `json:"invoice_number"`
	NumberOfInvoice    string `json:"number_of_invoice"`
	InvoiceDate        string `json:"invoice_date"`
	VendorName         string `json:"vendor_name"`
	FKVendor           string `json:"fk_vendor"`
	ItemDescription    string `json:"item_description"`
	UnitPrice          string `json:"unit_price"`
	VatUnitPrice       string `json:"vat_unit_price"`
	Quantity           string `json:"quantity"`
	PaymentTerm        string `json:"payment_term"`
	PaymentInstallment string `json:"payment_installment"`
	PaymentCenter      string `json:"payment_center"`
	PaymentType        string `json:"payment_type"`
	InvoiceTotal       string `json:"invoice_total"`
	VatInvoiceTotal    string `json:"vat_invoice_total"`

	IsAnotherItem string `json:"is_another_item"`

	Success string
	Error   string
}

// Validate validates the data of the purchase request sent by the user
func (purchaseRequestFormInput *PurchaseRequestFormInput) Validate() bool {

	purchaseRequestFormInput.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(purchaseRequestFormInput,
		validation.Field(&purchaseRequestFormInput.CostCenter, validation.Required),
		validation.Field(&purchaseRequestFormInput.Initiator, validation.Required),
		validation.Field(&purchaseRequestFormInput.PrType, validation.Required),
		validation.Field(&purchaseRequestFormInput.CostType, validation.Required),
		validation.Field(&purchaseRequestFormInput.CostCategory, validation.Required),
		validation.Field(&purchaseRequestFormInput.NumberOfInvoice, validation.Required),
		validation.Field(&purchaseRequestFormInput.InvoiceNumber, validation.Required, validation.Match(regexp.MustCompile(`^[0-9s\-]{0,500}$`)).Error("must look like 1234567-1234578, no space"), validation.Length(0, 500)),
		validation.Field(&purchaseRequestFormInput.InvoiceDate, validation.Required, validation.Date("1/2/2006")),
		validation.Field(&purchaseRequestFormInput.FKVendor, validation.Required),
		validation.Field(&purchaseRequestFormInput.ItemDescription, validation.Required),
		validation.Field(&purchaseRequestFormInput.UnitPrice, validation.Required, is.Float),
		validation.Field(&purchaseRequestFormInput.VatUnitPrice, validation.Required, is.Float),
		validation.Field(&purchaseRequestFormInput.Quantity, validation.Required, is.Int),
		validation.Field(&purchaseRequestFormInput.PaymentTerm, validation.Required),
		validation.Field(&purchaseRequestFormInput.PaymentInstallment, validation.Required),
		validation.Field(&purchaseRequestFormInput.PaymentCenter, validation.Required),
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
		`IDPurchaseRequest`:  purchaseRequestFormInput.IDPurchaseRequest,
		`CostCenter`:         purchaseRequestFormInput.CostCenter,
		`Initiator`:          purchaseRequestFormInput.Initiator,
		`PrType`:             purchaseRequestFormInput.PrType,
		`CostType`:           purchaseRequestFormInput.CostType,
		`CostCategory`:       purchaseRequestFormInput.CostCategory,
		`NumberOfInvoice`:    purchaseRequestFormInput.NumberOfInvoice,
		`InvoiceNumber`:      purchaseRequestFormInput.InvoiceNumber,
		`InvoiceDate`:        purchaseRequestFormInput.InvoiceDate,
		`VendorName`:         purchaseRequestFormInput.VendorName,
		`FKVendor`:           purchaseRequestFormInput.FKVendor,
		`ItemDescription`:    purchaseRequestFormInput.ItemDescription,
		`UnitPrice`:          purchaseRequestFormInput.UnitPrice,
		`VatUnitPrice`:       purchaseRequestFormInput.VatUnitPrice,
		`Quantity`:           purchaseRequestFormInput.Quantity,
		`PaymentTerm`:        purchaseRequestFormInput.PaymentTerm,
		`PaymentInstallment`: purchaseRequestFormInput.PaymentInstallment,
		`PaymentCenter`:      purchaseRequestFormInput.PaymentCenter,
		`PaymentType`:        purchaseRequestFormInput.PaymentType,
		`InvoiceTotal`:       purchaseRequestFormInput.InvoiceTotal,
		`VatInvoiceTotal`:    purchaseRequestFormInput.VatInvoiceTotal,

		`IsAnotherItem`: purchaseRequestFormInput.IsAnotherItem,

		`Success`: purchaseRequestFormInput.Success,
		`Error`:   purchaseRequestFormInput.Error,
	})
	handleErr(c, err)
}

func (purchaseRequestFormInput *PurchaseRequestFormInput) ChangeQuantity(newQuantity string) {
	purchaseRequestFormInput.Quantity = newQuantity
}

type CostCategory struct {
	IDCostCategory string `json:"id_cost_category"`
	Name           string `json:"name"`
	NameFa         string `json:"name_fa"`
}

type Vendor struct {
	IDVendor string `json:"id_vendor"`
	Name     string `json:"vendor_name"`
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
