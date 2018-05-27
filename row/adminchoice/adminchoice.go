package adminchoice

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

// AdminChoice represents one admin choice to accept or reject one purchase request
// it also includes potential error and a potential form to include into the web page used to collect admin choice
type AdminChoice struct {
	IDPurchaseRequest string `json:"id_purchase_request"`
	AcceptReject      string `json:"accept_reject"`

	Error string
}

// Validate validates the data of the admin choice sent by the user
func (adminChoice *AdminChoice) Validate() bool {

	adminChoice.Error = ""

	err := validation.ValidateStruct(adminChoice,
		validation.Field(&adminChoice.IDPurchaseRequest, validation.Required),
		validation.Field(&adminChoice.AcceptReject, validation.Required),
	)

	// add potential error text to adminChoice.Error
	if err != nil {
		adminChoice.Error = err.Error()
	}

	return adminChoice.Error == ""
}

// Render the web page itself given the html template and the adminChoice
func (adminChoice *AdminChoice) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the adminChoice
	err = tmpl.Execute(c.Writer, map[string]interface{}{`Error`: adminChoice.Error})
	handleErr(c, err)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
