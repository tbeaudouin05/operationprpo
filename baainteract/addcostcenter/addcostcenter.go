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
		) VALUES (@p1,@p2,@p3)`
	insertNewCostCenter, err := dbBaa.Prepare(insertNewCostCenterStr)

	res, err := insertNewCostCenter.Exec(
		costCenterFormInput.GFKDepartment,
		costCenterFormInput.FunctionCode,
		costCenterFormInput.FunctionName,
	)

	log.Println(res)

	return err
}

func CreateNewDepartment(departmentFormInput *costcenter.Department, dbBaa *sql.DB) error {

	insertNewDepartmentStr := `
	INSERT INTO baa_application.operation.department (
		gfk_location
		,department_code
		,name
		) VALUES (@p1,@p2,@p3)`
	insertNewDepartment, err := dbBaa.Prepare(insertNewDepartmentStr)

	res, err := insertNewDepartment.Exec(
		departmentFormInput.GFKLocation,
		departmentFormInput.DepartmentCode,
		departmentFormInput.DepartmentName,
	)

	log.Println(res)

	return err
}

func CreateNewLocation(locationFormInput *costcenter.Location, dbBaa *sql.DB) error {

	insertNewLocationStr := `
	INSERT INTO baa_application.operation.location (
		fk_division
		,location_code
		,name
		) VALUES (@p1,@p2,@p3)`
	insertNewLocation, err := dbBaa.Prepare(insertNewLocationStr)

	res, err := insertNewLocation.Exec(
		locationFormInput.FKDivision,
		locationFormInput.LocationCode,
		locationFormInput.LocationName,
	)

	log.Println(res)

	return err
}

// CHANGE TO ID COST CENTER: join with cost_center table to get name

func GetListOfFunctionName(dbBaa *sql.DB) string {

	// store query in a string
	query := `SELECT fu.name FROM baa_application.operation.func fu`

	var functionName, listOfFunctionName string
	rows, err := dbBaa.Query(query)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&functionName)
		checkError(err)
		listOfFunctionName += functionName + `,`
	}
	listOfFunctionName = listOfFunctionName[:len(listOfFunctionName)-1]
	return listOfFunctionName

}

func GetNextFunctionCode(dbBaa *sql.DB) string {

	// store query in a string
	query := `
	SELECT 
	CASE 
	WHEN COALESCE(MAX(fu.function_code)+1,0) < 10  THEN CONCAT('0',CAST(COALESCE(MAX(fu.function_code),0)+1 AS VARCHAR))
	ELSE  CAST(MAX(fu.function_code+1) AS VARCHAR) 
	END  next_function_code
  	FROM baa_application.operation.func fu`

	var nextFunctionCode string
	err := dbBaa.QueryRow(query).Scan(&nextFunctionCode)
	checkError(err)
	return nextFunctionCode

}

func GetListOfDepartmentName(dbBaa *sql.DB) string {

	// store query in a string
	query := `SELECT de.name FROM baa_application.operation.department de`

	var departmentName, listOfDepartmentName string
	rows, err := dbBaa.Query(query)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&departmentName)
		checkError(err)
		listOfDepartmentName += departmentName + `,`
	}
	listOfDepartmentName = listOfDepartmentName[:len(listOfDepartmentName)-1]
	return listOfDepartmentName

}

func GetNextDepartmentCode(dbBaa *sql.DB) string {

	// store query in a string
	query := `
	SELECT 
  	CASE 
  	WHEN COALESCE(MAX(d.department_code)+1,0) < 10  THEN CONCAT('0',CAST(COALESCE(MAX(d.department_code),0)+1 AS VARCHAR))
  	ELSE  CAST(MAX(d.department_code+1) AS VARCHAR) 
  	END  next_department_code
	FROM baa_application.operation.department d`

	var nextDepartmentCode string
	err := dbBaa.QueryRow(query).Scan(&nextDepartmentCode)
	checkError(err)
	return nextDepartmentCode

}

func GetListOfLocationName(dbBaa *sql.DB) string {

	// store query in a string
	query := `SELECT lo.name FROM baa_application.operation.location lo`

	var locationName, listOfLocationName string
	rows, err := dbBaa.Query(query)
	checkError(err)
	for rows.Next() {
		err := rows.Scan(&locationName)
		checkError(err)
		listOfLocationName += locationName + `,`
	}
	listOfLocationName = listOfLocationName[:len(listOfLocationName)-1]
	return listOfLocationName

}

func GetNextLocationCode(dbBaa *sql.DB) string {

	// store query in a string
	query := `
	SELECT 
	CASE 
	WHEN COALESCE(MAX(l.location_code)+1,0) < 10 THEN CONCAT('0',CAST(COALESCE(MAX(l.location_code),0)+1 AS VARCHAR))
	ELSE  CAST(MAX(l.location_code+1) AS VARCHAR) 
	END  next_location_code
  	FROM baa_application.operation.location l`

	var nextLocationCode string
	err := dbBaa.QueryRow(query).Scan(&nextLocationCode)
	checkError(err)
	return nextLocationCode

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
