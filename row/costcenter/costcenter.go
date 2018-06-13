package costcenter

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/gin-gonic/gin"
)

// CostCenter stores necessary information to register a new CostCenter
type CostCenter struct {
	GIDFunction       string `json:"gid_function"`
	GFKDepartment     string `json:"gfk_department"`
	DepartmentName    string `json:"department_name"`
	FunctionCode      string `json:"function_code"`
	FunctionName      string `json:"function_name"`
	FunctionTag       string `json:"function_tag"`
	FunctionTagCode   string `json:"function_tag_code"`
	FunctionNameFarsi string `json:"function_name_farsi"`

	Error   string
	Success string
}

// ValidateCostCenter validates the data of the costCenter
func (costCenter *CostCenter) ValidateCostCenter(listOfExistingGIDFunction []string) bool {

	costCenter.Error = ""

	// define validation of each field of the costCenter
	err := validation.ValidateStruct(costCenter,
		validation.Field(&costCenter.GIDFunction, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{9}$")), notInArray(listOfExistingGIDFunction)),
		validation.Field(&costCenter.GFKDepartment, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{6}$"))),
		validation.Field(&costCenter.FunctionCode, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{3}$"))),
		validation.Field(&costCenter.FunctionName, validation.Required, validation.Length(0, 100)),
		validation.Field(&costCenter.FunctionTag, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{4}$"))),
		validation.Field(&costCenter.FunctionTagCode, validation.Match(regexp.MustCompile("^[0-9]{3}$"))),
		validation.Field(&costCenter.FunctionNameFarsi, validation.Length(0, 100)),
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
		`GIDFunction`:       costCenter.GIDFunction,
		`GFKDepartment`:     costCenter.GFKDepartment,
		`FunctionCode`:      costCenter.FunctionCode,
		`FunctionName`:      costCenter.FunctionName,
		`FunctionTag`:       costCenter.FunctionTag,
		`FunctionTagCode`:   costCenter.FunctionTagCode,
		`FunctionNameFarsi`: costCenter.FunctionNameFarsi,

		`Error`:   costCenter.Error,
		`Success`: costCenter.Success,
	})
	handleErr(c, err)
}

// Department ------------------------------------------------------------------------------------------------------

// Department stores necessary information to register a new Department
type Department struct {
	GIDDepartment       string `json:"gid_department"`
	GFKLocation         string `json:"gfk_location"`
	DepartmentShortName string `json:"department_short_name"`
	DepartmentCode      string `json:"department_code"`
	DepartmentName      string `json:"department_name"`
	DepartmentTag       string `json:"department_tag"`
	DepartmentTagCode   string `json:"department_tag_code"`
	DepartmentNameFarsi string `json:"department_name_farsi"`

	Error   string
	Success string
}

// ValidateDepartment validates the data of the department
func (department *Department) ValidateDepartment(listOfExistingGIDDepartment []string) bool {

	department.Error = ""

	// define validation of each field of the department
	err := validation.ValidateStruct(department,
		validation.Field(&department.GIDDepartment, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{6}$")), notInArray(listOfExistingGIDDepartment)),
		validation.Field(&department.GFKLocation, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{4}$"))),
		validation.Field(&department.DepartmentCode, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{2}$"))),
		validation.Field(&department.DepartmentName, validation.Required, validation.Length(0, 100)),
		validation.Field(&department.DepartmentTag, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{3}$"))),
		validation.Field(&department.DepartmentTagCode, validation.Match(regexp.MustCompile("^[0-9]{2}$"))),
		validation.Field(&department.DepartmentNameFarsi, validation.Length(0, 100)),
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
		`GIDDepartment`:       department.GIDDepartment,
		`GFKLocation`:         department.GFKLocation,
		`DepartmentCode`:      department.DepartmentCode,
		`DepartmentName`:      department.DepartmentName,
		`DepartmentTag`:       department.DepartmentTag,
		`DepartmentTagCode`:   department.DepartmentTagCode,
		`DepartmentNameFarsi`: department.DepartmentNameFarsi,

		`Error`:   department.Error,
		`Success`: department.Success,
	})
	handleErr(c, err)
}

// Division---------------------------------------------------------------------------------------

// Location stores necessary information to register a new Location
type Location struct {
	GIDLocation       string `json:"gid_location"`
	FKDivision        string `json:"gfk_location"`
	LocationShortName string `json:"location_short_name"`
	LocationCode      string `json:"location_code"`
	LocationName      string `json:"location_name"`
	LocationTag       string `json:"location_tag"`
	LocationTagCode   string `json:"location_tag_code"`
	LocationNameFarsi string `json:"location_name_farsi"`

	Error   string
	Success string
}

// ValidateLocation validates the data of the location
func (location *Location) ValidateLocation(listOfExistingGIDLocation []string) bool {

	location.Error = ""

	// define validation of each field of the location
	err := validation.ValidateStruct(location,
		validation.Field(&location.GIDLocation, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{4}$")), notInArray(listOfExistingGIDLocation)),
		validation.Field(&location.FKDivision, validation.Required, validation.Match(regexp.MustCompile("^6[0-9]{1}$"))),
		validation.Field(&location.LocationCode, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{2}$"))),
		validation.Field(&location.LocationName, validation.Required, validation.Length(0, 100)),
		validation.Field(&location.LocationTag, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{3}$"))),
		validation.Field(&location.LocationTagCode, validation.Match(regexp.MustCompile("^[0-9]{2}$"))),
		validation.Field(&location.LocationNameFarsi, validation.Length(0, 100)),
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
		`GIDLocation`:       location.GIDLocation,
		`FKDivision`:        location.FKDivision,
		`LocationCode`:      location.LocationCode,
		`LocationName`:      location.LocationName,
		`LocationTag`:       location.LocationTag,
		`LocationTagCode`:   location.LocationTagCode,
		`LocationNameFarsi`: location.LocationNameFarsi,

		`Error`:   location.Error,
		`Success`: location.Success,
	})
	handleErr(c, err)
}

func notInArray(values []string) *notInArrayRule {
	return &notInArrayRule{
		elements: values,
		message:  "the cost center, department or location already exists!",
	}
}

type notInArrayRule struct {
	elements []string
	message  string
}

// Validate checks if the given value is valid or not.
func (r *notInArrayRule) Validate(value interface{}) error {

	for _, e := range r.elements {
		if e == value {
			return errors.New(r.message)
		}
	}
	return nil
}

// Error sets the error message for the rule.
func (r *notInArrayRule) Error(message string) *notInArrayRule {
	r.message = message
	return r
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
