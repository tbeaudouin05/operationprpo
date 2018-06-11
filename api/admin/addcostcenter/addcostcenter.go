package addcostcenter

import (
	"fmt"
	"net/http"

	"github.com/thomas-bamilo/operation/operationprpo/row/costcenter"

	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/operation/operationprpo/api/oauth/authorize"
	"github.com/thomas-bamilo/operation/operationprpo/baainteract"
	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"
	"github.com/thomas-bamilo/sql/connectdb"
)

var user useraccess.User

// Start loads the purchase request form web page - GET request
func Start(c *gin.Context) {

	authorize.AuthorizeCcAdmin(c, &user)

	// only pass form as addUserFormInput since we only want a blank form at start
	costCenterFormInput := &costcenter.CostCenter{}

	// render the web page itself given the html template and the addUserFormInput
	costCenterFormInput.Render(c, `template/admin/addcostcenter/addcostcenter.html`)
}

// AnswerCostCenterForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerCostCenterForm(c *gin.Context) {

	authorize.AuthorizeCcAdmin(c, &user)

	r := c.Request

	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	costCenterFormInput := &costcenter.CostCenter{
		GFKDepartment: r.FormValue(`GFKDepartment`),
		FunctionCode:  baainteract.GetNextFunctionCode(dbBaa),
		FunctionName:  r.FormValue(`FunctionName`),
	}

	listOfFunctionName := baainteract.GetListOfFunctionName(dbBaa)

	if costCenterFormInput.Validate(listOfFunctionName) == false {
		costCenterFormInput.Render(c, `template/admin/addcostcenter.html`)
		return
	}

	err := baainteract.CreateNewCostCenter(costCenterFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to addcostcenterconfirmation web page
	http.Redirect(c.Writer, r, `/admin/addcostcenterconfirmation`, http.StatusSeeOther)
}

// AnswerDepartmentForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerDepartmentForm(c *gin.Context) {

	authorize.AuthorizeCcAdmin(c, &user)

	r := c.Request

	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	departmentFormInput := &costcenter.Department{
		GFKLocation:    r.FormValue(`GFKLocation`),
		DepartmentCode: baainteract.GetNextDepartmentCode(dbBaa),
		DepartmentName: r.FormValue(`DepartmentName`),
	}

	listOfDepartmentName := baainteract.GetListOfDepartmentName(dbBaa)

	if departmentFormInput.Validate(listOfDepartmentName) == false {
		departmentFormInput.Render(c, `template/admin/adddepartment.html`)
		return
	}

	err := baainteract.CreateNewDepartment(departmentFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to addcostcenterconfirmation web page
	http.Redirect(c.Writer, r, `/admin/addcostcenterconfirmation`, http.StatusSeeOther)
}

// AnswerLocationForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerLocationForm(c *gin.Context) {

	authorize.AuthorizeCcAdmin(c, &user)

	r := c.Request

	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	locationFormInput := &costcenter.Location{
		FKDivision:   r.FormValue(`FKDivision`),
		LocationCode: baainteract.GetNextLocationCode(dbBaa),
		LocationName: r.FormValue(`LocationName`),
	}

	listOfLocationName := baainteract.GetListOfLocationName(dbBaa)

	if locationFormInput.Validate(listOfLocationName) == false {
		locationFormInput.Render(c, `template/admin/addlocation.html`)
		return
	}

	err := baainteract.CreateNewLocation(locationFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to addcostcenterconfirmation web page
	http.Redirect(c.Writer, r, `/admin/addcostcenterconfirmation`, http.StatusSeeOther)
}

// ConfirmForm loads the purchase request addcostcenterconfirmation web page - GET request
func ConfirmForm(c *gin.Context) {

	costCenterFormInput := &costcenter.CostCenter{}

	// render addcostcenterconfirmation web page
	costCenterFormInput.Render(c, `template/admin/addcostcenter/addcostcenterconfirmation.html`)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
