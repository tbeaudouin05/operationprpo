package addcostcenter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thomas-bamilo/operation/operationprpo/row/costcenter"

	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/operation/operationprpo/api/oauth/authorize"
	cc_baainteract "github.com/thomas-bamilo/operation/operationprpo/baainteract/addcostcenter"
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

	functionCode := r.FormValue(`FunctionCode`)
	gFKDepartment := r.FormValue(`DepartmentName`)
	gIDFunction := gFKDepartment + functionCode

	costCenterFormInput := &costcenter.CostCenter{
		GIDFunction:       gIDFunction,
		GFKDepartment:     gFKDepartment,
		FunctionCode:      functionCode,
		FunctionName:      r.FormValue(`FunctionName`),
		FunctionTag:       r.FormValue(`FunctionTag`),
		FunctionTagCode:   r.FormValue(`FunctionTagCode`),
		FunctionNameFarsi: r.FormValue(`FunctionNameFarsi`),
	}

	existingGIDFunction := cc_baainteract.GetExistingGIDFunction(dbBaa)

	if costCenterFormInput.ValidateCostCenter(existingGIDFunction) == false {
		costCenterFormInput.Render(c, `template/admin/addcostcenter/addcostcenter.html`)
		return
	}

	err := cc_baainteract.CreateNewCostCenter(costCenterFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to addcostcenterconfirmation web page
	http.Redirect(c.Writer, r, `/admin/addcostcenterconfirmation`, http.StatusSeeOther)
}

// StartExistingFunctionName - GET request
func StartExistingFunctionName(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	// GetExistingFunctionName queries baa_application.operation.purchase_request to return all pending purchase requests
	ExistingFunctionNameTable := cc_baainteract.GetExistingFunctionName(dbBaa)

	//Convert the `ExistingFunctionNameTable` variable to json
	ExistingFunctionNameTableByte, err := json.Marshal(ExistingFunctionNameTable)
	handleErr(c, err)

	// If all goes well, write the JSON list of ExistingFunctionNameTable to the response
	c.Writer.Write(ExistingFunctionNameTableByte)
}

// StartExistingCostCenterName - GET request
func StartExistingCostCenterName(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	// GetExistingCostCenterName queries baa_application.operation.purchase_request to return all pending purchase requests
	ExistingCostCenterNameTable := cc_baainteract.GetExistingCostCenterName(dbBaa)

	//Convert the `ExistingCostCenterNameTable` variable to json
	ExistingCostCenterNameTableByte, err := json.Marshal(ExistingCostCenterNameTable)
	handleErr(c, err)

	// If all goes well, write the JSON list of ExistingCostCenterNameTable to the response
	c.Writer.Write(ExistingCostCenterNameTableByte)
}

// Department -------------------------------------------------------------------------------------------------------------------

// AnswerDepartmentForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerDepartmentForm(c *gin.Context) {

	authorize.AuthorizeCcAdmin(c, &user)

	r := c.Request

	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	departmentFormInput := &costcenter.Department{
		GIDDepartment:       r.FormValue(`LocationName`) + r.FormValue(`DepartmentCode`),
		GFKLocation:         r.FormValue(`LocationName`),
		DepartmentCode:      r.FormValue(`DepartmentCode`),
		DepartmentName:      r.FormValue(`DepartmentName`),
		DepartmentTag:       r.FormValue(`DepartmentTag`),
		DepartmentTagCode:   r.FormValue(`DepartmentTagCode`),
		DepartmentNameFarsi: r.FormValue(`DepartmentNameFarsi`),
	}

	existingGIDDepartment := cc_baainteract.GetExistingGIDDepartment(dbBaa)

	if departmentFormInput.ValidateDepartment(existingGIDDepartment) == false {
		departmentFormInput.Render(c, `template/admin/addcostcenter/addcostcenter.html`)
		return
	}

	err := cc_baainteract.CreateNewDepartment(departmentFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to addcostcenterconfirmation web page
	http.Redirect(c.Writer, r, `/admin/addcostcenterconfirmation`, http.StatusSeeOther)
}

// StartExistingDepartmentName - GET request
func StartExistingDepartmentName(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	// GetExistingDepartmentName queries baa_application.operation.purchase_request to return all pending purchase requests
	ExistingDepartmentNameTable := cc_baainteract.GetExistingDepartmentName(dbBaa)

	//Convert the `ExistingDepartmentNameTable` variable to json
	ExistingDepartmentNameTableByte, err := json.Marshal(ExistingDepartmentNameTable)
	handleErr(c, err)

	// If all goes well, write the JSON list of ExistingDepartmentNameTable to the response
	c.Writer.Write(ExistingDepartmentNameTableByte)
}

// Location -------------------------------------------------------------------------------------------------------------

// AnswerLocationForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerLocationForm(c *gin.Context) {

	authorize.AuthorizeCcAdmin(c, &user)

	r := c.Request

	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	locationFormInput := &costcenter.Location{
		FKDivision:        "69",
		LocationCode:      r.FormValue(`LocationCode`),
		LocationName:      r.FormValue(`LocationName`),
		LocationTag:       r.FormValue(`LocationTag`),
		LocationTagCode:   r.FormValue(`LocationTagCode`),
		LocationNameFarsi: r.FormValue(`LocationNameFarsi`),
	}

	existingGIDLocation := cc_baainteract.GetExistingGIDLocation(dbBaa)

	if locationFormInput.ValidateLocation(existingGIDLocation) == false {
		locationFormInput.Render(c, `template/admin/addcostcenter/addcostcenter.html`)
		return
	}

	err := cc_baainteract.CreateNewLocation(locationFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to addcostcenterconfirmation web page
	http.Redirect(c.Writer, r, `/admin/addcostcenterconfirmation`, http.StatusSeeOther)
}

// StartExistingLocationName - GET request
func StartExistingLocationName(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	// GetExistingLocationName queries baa_application.operation.purchase_request to return all pending purchase requests
	ExistingLocationNameTable := cc_baainteract.GetExistingLocationName(dbBaa)

	//Convert the `ExistingLocationNameTable` variable to json
	ExistingLocationNameTableByte, err := json.Marshal(ExistingLocationNameTable)
	handleErr(c, err)

	// If all goes well, write the JSON list of ExistingLocationNameTable to the response
	c.Writer.Write(ExistingLocationNameTableByte)
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
