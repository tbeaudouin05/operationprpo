package useraccess

import (
	"fmt"
	"html/template"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/gin-gonic/gin"
)

// User is a retrieved and authenticated user.
type User struct {
	Sub           string `json:"sub"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
	HD            string `json:"hd"`
	LoginURL      string `json:"login_url"`

	IDUser           string `json:"id_user"`
	Email            string `json:"email"`
	Name             string `json:"name"`
	Access           string `json:"access"`
	LocationAccess   string `json:"location_access"`
	DepartmentAccess string `json:"department_access"`

	// to allow user access to department
	GFKDepartment  string `json:"gfk_department"`
	DepartmentName string `json:"department_name"`

	// to allow user access to location
	GFKLocation  string `json:"gfk_location"`
	LocationName string `json:"location_name"`

	Error   string
	Success string
}

// Validate validates the Email and Name of the User
func (user *User) Validate() bool {

	user.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Name, validation.Required),
	)

	// add potential error text to user.Error
	if err != nil {
		user.Error = err.Error()
	}

	// return true if no error, false otherwise
	return user.Error == ""
}

// ValidateDepartmentAccess validates the DepartmentAccess of the user
func (user *User) ValidateDepartmentAccess() bool {

	user.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(user,
		validation.Field(&user.IDUser, validation.Required),
		validation.Field(&user.GFKDepartment, validation.Required),
	)

	// add potential error text to user.Error
	if err != nil {
		user.Error = err.Error()
	}

	// return true if no error, false otherwise
	return user.Error == ""
}

// ValidateLocationAccess validates the LocationAccess of the user
func (user *User) ValidateLocationAccess() bool {

	user.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(user,
		validation.Field(&user.IDUser, validation.Required),
		validation.Field(&user.GFKLocation, validation.Required),
	)

	// add potential error text to user.Error
	if err != nil {
		user.Error = err.Error()
	}

	// return true if no error, false otherwise
	return user.Error == ""
}

// Render the web page itself given the html template and the user
func (user *User) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the user
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`EmailVerified`: user.EmailVerified,
		`HD`:            user.HD,
		`Access`:        user.Access,
		`LoginURL`:      user.LoginURL,

		`IDUser`: user.IDUser,
		`Email`:  user.Email,
		`Name`:   user.Name,

		`GFKDepartment`:  user.GFKDepartment,
		`DepartmentName`: user.DepartmentName,

		`GFKLocation`:  user.GFKLocation,
		`LocationName`: user.LocationName,

		`Error`:   user.Error,
		`Success`: user.Success,
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
