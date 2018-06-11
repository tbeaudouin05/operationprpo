package costcenter

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/gin-gonic/gin"
)

// CostCenter stores necessary information to register a new CostCenter
type CostCenter struct {
	GIDFunction   string `json:"gid_function"`
	GFKDepartment string `json:"gfk_department"`
	FunctionCode  string `json:"function_code"`
	FunctionName  string `json:"function_name"`

	Error   string
	Success string
}

// Validate validates the data of the costCenter
func (costCenter *CostCenter) Validate(listOfFunctionName string) bool {

	costCenter.Error = ""

	// define validation of each field of the costCenter
	err := validation.ValidateStruct(costCenter,
		validation.Field(&costCenter.GFKDepartment, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{6}$"))),
		validation.Field(&costCenter.FunctionCode, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{3}$")), validation.NotIn(listOfFunctionName)),
		validation.Field(&costCenter.FunctionName, validation.Required, validation.NotIn(listOfFunctionName)),
	)

	// add potential error text to costCenter.Error
	if err != nil {
		costCenter.Error = err.Error()
	}

	// return true if no error, false otherwise
	return costCenter.Error == ""
}

// Render the web page itself given the html template and the costCenter
func (costCenter *CostCenter) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the costCenter
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`GIDFunction`:   costCenter.GIDFunction,
		`GFKDepartment`: costCenter.GFKDepartment,
		`FunctionCode`:  costCenter.FunctionCode,
		`FunctionName`:  costCenter.FunctionName,

		`Error`:   costCenter.Error,
		`Success`: costCenter.Success,
	})
	handleErr(c, err)
}

// Department -----------------------------------------------------------

// Department stores necessary information to register a new Department
type Department struct {
	GFKLocation    string `json:"gfk_location"`
	DepartmentCode string `json:"department_code"`
	DepartmentName string `json:"department_name"`

	Error   string
	Success string
}

// Validate validates the data of the department
func (department *Department) Validate(listOfDepartmentName string) bool {

	department.Error = ""

	// define validation of each field of the department
	err := validation.ValidateStruct(department,
		validation.Field(&department.GFKLocation, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{4}$"))),
		validation.Field(&department.DepartmentCode, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{2}$"))),
		validation.Field(&department.DepartmentName, validation.Required, validation.NotIn(listOfDepartmentName)),
	)

	// add potential error text to department.Error
	if err != nil {
		department.Error = err.Error()
	}

	// return true if no error, false otherwise
	return department.Error == ""
}

// Render the web page itself given the html template and the department
func (department *Department) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the department
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`GFKLocation`:    department.GFKLocation,
		`DepartmentCode`: department.DepartmentCode,
		`DepartmentName`: department.DepartmentName,

		`Error`:   department.Error,
		`Success`: department.Success,
	})
	handleErr(c, err)
}

// Location -----------------------------------------------------------

// Location stores necessary information to register a new Location
type Location struct {
	FKDivision   string `json:"fk_division"`
	LocationCode string `json:"location_code"`
	LocationName string `json:"location_name"`

	Error   string
	Success string
}

// Validate validates the data of the location
func (location *Location) Validate(listOfLocationName string) bool {

	location.Error = ""

	// define validation of each field of the location
	err := validation.ValidateStruct(location,
		validation.Field(&location.FKDivision, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{2}$"))),
		validation.Field(&location.LocationCode, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{2}$"))),
		validation.Field(&location.LocationName, validation.Required, validation.NotIn(listOfLocationName)),
	)

	// add potential error text to location.Error
	if err != nil {
		location.Error = err.Error()
	}

	// return true if no error, false otherwise
	return location.Error == ""
}

// Render the web page itself given the html template and the location
func (location *Location) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the location
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`FKDivision`:   location.FKDivision,
		`LocationCode`: location.LocationCode,
		`LocationName`: location.LocationName,

		`Error`:   location.Error,
		`Success`: location.Success,
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
