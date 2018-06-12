package adduser

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/operation/operationprpo/api/oauth/authorize"
	us_baainteract "github.com/thomas-bamilo/operation/operationprpo/baainteract/adduser"
	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"
	"github.com/thomas-bamilo/sql/connectdb"
)

var user useraccess.User

// Start loads the purchase request form web page - GET request
func Start(c *gin.Context) {

	authorize.AuthorizePrAdmin(c, &user)

	// only pass form as addUserFormInput since we only want a blank form at start
	addUserFormInput := &useraccess.User{}

	// render the web page itself given the html template and the addUserFormInput
	addUserFormInput.Render(c, `template/admin/adduser/adduser.html`)
}

// AnswerForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerForm(c *gin.Context) {

	authorize.AuthorizePrAdmin(c, &user)

	r := c.Request

	// pass all the form values input by the user
	// since we want to validate these values and upload them to database
	// in case validation fails, we also want to return these values to the form for good user experience
	addUserFormInput := &useraccess.User{
		Email: r.FormValue(`email`),
		Name:  r.FormValue(`name`),
	}

	// Validate validates the addUserFormInput form user inputs
	// if validation fails, reload the purchase request form page with the initial user inputs and error messages
	if addUserFormInput.Validate() == false {
		addUserFormInput.Render(c, `template/admin/adduser.html`)
		return
	}

	// LoadToDb uploads addUserFormInput to database
	dbBaa := connectdb.ConnectToBaa()
	err := us_baainteract.CreateNewUser(addUserFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to adduserconfirmation web page
	http.Redirect(c.Writer, r, `/admin/adduserconfirmation`, http.StatusSeeOther)
}

// AnswerDepartmentAccessForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerDepartmentAccessForm(c *gin.Context) {

	authorize.AuthorizePrAdmin(c, &user)

	r := c.Request

	departmentAccessFormInput := &useraccess.User{
		IDUser:        r.FormValue(`emaildepartment`),
		GFKDepartment: r.FormValue(`departmentaccess`),
	}

	if departmentAccessFormInput.ValidateDepartmentAccess() == false {
		departmentAccessFormInput.Render(c, `template/admin/adduser.html`)
		return
	}

	// LoadToDb uploads to database
	dbBaa := connectdb.ConnectToBaa()
	err := us_baainteract.AddUserDepartmentAccess(departmentAccessFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to adduserconfirmation web page
	http.Redirect(c.Writer, r, `/admin/adduserconfirmation`, http.StatusSeeOther)
}

// AnswerLocationAccessForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerLocationAccessForm(c *gin.Context) {

	authorize.AuthorizePrAdmin(c, &user)

	r := c.Request

	userLocationAccessFormInput := &useraccess.User{
		IDUser:      r.FormValue(`emaillocation`),
		GFKLocation: r.FormValue(`locationaccess`),
	}

	if userLocationAccessFormInput.ValidateLocationAccess() == false {
		userLocationAccessFormInput.Render(c, `template/admin/adduser.html`)
		return
	}

	// LoadToDb uploads to database
	dbBaa := connectdb.ConnectToBaa()
	err := us_baainteract.AddUserLocationAccess(userLocationAccessFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to adduserconfirmation web page
	http.Redirect(c.Writer, r, `/admin/adduserconfirmation`, http.StatusSeeOther)
}

// ConfirmForm loads the purchase request adduserconfirmation web page - GET request
func ConfirmForm(c *gin.Context) {

	addUserFormInput := &useraccess.User{}

	// render adduserconfirmation web page
	addUserFormInput.Render(c, `template/admin/adduser/adduserconfirmation.html`)
}

//  - GET request
func StartIDEmail(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	userIDEmailTable := us_baainteract.GetUserIDEmail(dbBaa)

	//Convert to json
	userIDEmailTableByte, err := json.Marshal(userIDEmailTable)
	handleErr(c, err)

	// If all goes well, write the JSON to the response
	c.Writer.Write(userIDEmailTableByte)
}

//  - GET request
func StartDepartmentAccess(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	departmentAccessTable := us_baainteract.GetDepartmentAccess(dbBaa)

	//Convert to json
	departmentAccessTableByte, err := json.Marshal(departmentAccessTable)
	handleErr(c, err)

	// If all goes well, write the JSON to the response
	c.Writer.Write(departmentAccessTableByte)
}

//  - GET request
func StartLocationAccess(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	locationAccessTable := us_baainteract.GetLocationAccess(dbBaa)

	//Convert to json
	locationAccessTableByte, err := json.Marshal(locationAccessTable)
	handleErr(c, err)

	// If all goes well, write the JSON to the response
	c.Writer.Write(locationAccessTableByte)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
