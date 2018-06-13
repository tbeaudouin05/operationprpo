package addcostcenter

import (
	"database/sql"
	"log"

	"github.com/thomas-bamilo/operation/operationprpo/row/costcenter"
)

// CHANGE TO ID COST CENTER

func CreateNewCostCenter(costCenterFormInput *costcenter.CostCenter, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertNewCostCenterStr := `
	INSERT INTO baa_application.operation.func (
		gfk_department
		,function_code
		,name
		,tag
		,tag_code
		,name_fa
		) VALUES (@p1,@p2,@p3,@p4,@p5,@p6)`
	insertNewCostCenter, err := dbBaa.Prepare(insertNewCostCenterStr)

	res, err := insertNewCostCenter.Exec(
		costCenterFormInput.GFKDepartment,
		costCenterFormInput.FunctionCode,
		costCenterFormInput.FunctionName,
		costCenterFormInput.FunctionTag,
		costCenterFormInput.FunctionTagCode,
		costCenterFormInput.FunctionNameFarsi,
	)

	log.Println(res)

	return err
}

func GetExistingGIDFunction(dbBaa *sql.DB) []string {

	// store query in a string
	query := `SELECT fu.gid_function FROM baa_application.operation.func fu`

	var gIDFunction string
	var listOfGIDFunction []string
	rows, err := dbBaa.Query(query)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&gIDFunction)
		checkError(err)
		listOfGIDFunction = append(listOfGIDFunction, gIDFunction)
	}
	listOfGIDFunction = listOfGIDFunction[:len(listOfGIDFunction)-1]
	return listOfGIDFunction

}

// GetExistingFunctionName
func GetExistingFunctionName(dbBaa *sql.DB) []*costcenter.CostCenter {
	rows, err := dbBaa.Query(`
		SELECT

		f.function_code
		,f.name
		
	  FROM baa_application.operation.func f
	  GROUP BY f.function_code, f.name`)

	checkError(err)
	defer rows.Close()

	// Create the data structure that is returned from the function.
	costCenterTable := []*costcenter.CostCenter{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a costCenter,
		costCenter := &costcenter.CostCenter{}
		// Populate the attributes of the costCenter,
		// and return incase of an error
		err := rows.Scan(
			&costCenter.FunctionCode,
			&costCenter.FunctionName,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		costCenterTable = append(costCenterTable, costCenter)
	}

	return costCenterTable

}

// GetExistingCostCenterName
func GetExistingCostCenterName(dbBaa *sql.DB) []*costcenter.CostCenter {
	rows, err := dbBaa.Query(`
		SELECT 

		ccv.gid_cost_center
		,ccv.cost_center_name
		
	  FROM baa_application.operation.cost_center_view ccv`)

	checkError(err)
	defer rows.Close()

	// Create the data structure that is returned from the function.
	costCenterTable := []*costcenter.CostCenter{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a costCenter,
		costCenter := &costcenter.CostCenter{}
		// Populate the attributes of the costCenter,
		// and return incase of an error
		err := rows.Scan(
			&costCenter.GIDFunction,
			&costCenter.FunctionName,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		costCenterTable = append(costCenterTable, costCenter)
	}

	return costCenterTable

}

// Department ----------------------------------------------------------------------------------------------------

func CreateNewDepartment(departmentFormInput *costcenter.Department, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertNewDepartmentStr := `
	INSERT INTO baa_application.operation.department (
		gfk_location
		,department_code
		,name
		,tag
		,tag_code
		,name_fa
		) VALUES (@p1,@p2,@p3,@p4,@p5,@p6)`
	insertNewDepartment, err := dbBaa.Prepare(insertNewDepartmentStr)

	res, err := insertNewDepartment.Exec(
		departmentFormInput.GFKLocation,
		departmentFormInput.DepartmentCode,
		departmentFormInput.DepartmentName,
		departmentFormInput.DepartmentTag,
		departmentFormInput.DepartmentTagCode,
		departmentFormInput.DepartmentNameFarsi,
	)

	log.Println(res)

	return err
}

func GetExistingGIDDepartment(dbBaa *sql.DB) []string {

	// store query in a string
	query := `SELECT de.gid_department FROM baa_application.operation.department de`

	var gIDDepartment string
	var listOfGIDDepartment []string
	rows, err := dbBaa.Query(query)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&gIDDepartment)
		checkError(err)
		listOfGIDDepartment = append(listOfGIDDepartment, gIDDepartment)
	}
	listOfGIDDepartment = listOfGIDDepartment[:len(listOfGIDDepartment)-1]
	return listOfGIDDepartment

}

// GetExistingDepartmentName
func GetExistingDepartmentName(dbBaa *sql.DB) []*costcenter.Department {
	rows, err := dbBaa.Query(`
		SELECT 

		ccv.gfk_department
		,ccv.department_name
		
	  FROM baa_application.operation.cost_center_view ccv

	  GROUP BY ccv.gfk_department, ccv.department_name`)

	checkError(err)
	defer rows.Close()

	// Create the data structure that is returned from the department.
	departmentTable := []*costcenter.Department{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a department,
		department := &costcenter.Department{}
		// Populate the attributes of the department,
		// and return incase of an error
		err := rows.Scan(
			&department.GIDDepartment,
			&department.DepartmentName,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		departmentTable = append(departmentTable, department)
	}

	return departmentTable

}

// Location --------------------------------------------------------------------------------------------------

func CreateNewLocation(locationFormInput *costcenter.Location, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertNewLocationStr := `
	INSERT INTO baa_application.operation.division (
		fk_division
		,location_code
		,name
		,tag
		,tag_code
		,name_fa
		) VALUES (@p1,@p2,@p3,@p4,@p5,@p6)`
	insertNewLocation, err := dbBaa.Prepare(insertNewLocationStr)

	res, err := insertNewLocation.Exec(
		locationFormInput.FKDivision,
		locationFormInput.LocationCode,
		locationFormInput.LocationName,
		locationFormInput.LocationTag,
		locationFormInput.LocationTagCode,
		locationFormInput.LocationNameFarsi,
	)

	log.Println(res)

	return err
}

func GetExistingGIDLocation(dbBaa *sql.DB) []string {

	// store query in a string
	query := `SELECT lo.gid_location FROM baa_application.operation.location lo`

	var gIDLocation string
	var listOfGIDLocation []string
	rows, err := dbBaa.Query(query)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&gIDLocation)
		checkError(err)
		listOfGIDLocation = append(listOfGIDLocation, gIDLocation)
	}
	listOfGIDLocation = listOfGIDLocation[:len(listOfGIDLocation)-1]
	return listOfGIDLocation

}

// GetExistingLocationName
func GetExistingLocationName(dbBaa *sql.DB) []*costcenter.Location {
	rows, err := dbBaa.Query(`
		SELECT 

		ccv.gfk_location
		,ccv.location_name
		
	  FROM baa_application.operation.cost_center_view ccv
	  
	  GROUP BY ccv.gfk_location, ccv.location_name`)

	checkError(err)
	defer rows.Close()

	// Create the data structure that is returned from the location.
	locationTable := []*costcenter.Location{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a location,
		location := &costcenter.Location{}
		// Populate the attributes of the location,
		// and return incase of an error
		err := rows.Scan(
			&location.GIDLocation,
			&location.LocationName,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		locationTable = append(locationTable, location)
	}

	return locationTable

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
