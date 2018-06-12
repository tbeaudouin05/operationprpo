package adduser

import (
	"database/sql"
	"log"

	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"
)

func CreateNewUser(userFormInput *useraccess.User, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertNewUserStr := `
INSERT INTO baa_application.operation.pr_user (
	email 
	,name
	,access) 
VALUES (@p1,@p2,'user')`
	insertNewUser, err := dbBaa.Prepare(insertNewUserStr)

	res, err := insertNewUser.Exec(
		userFormInput.Email,
		userFormInput.Name,
	)

	log.Println(res)

	return err
}

func AddUserDepartmentAccess(userDepartmentAccessFormInput *useraccess.User, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertNewUserStr := `
	INSERT INTO  baa_application.operation.pr_department_access (
		fk_user
		,gfk_department)
	VALUES (@p1, @p2)
`
	insertNewUser, err := dbBaa.Prepare(insertNewUserStr)

	res, err := insertNewUser.Exec(
		userDepartmentAccessFormInput.IDUser,
		userDepartmentAccessFormInput.GFKDepartment,
	)

	log.Println(res)

	return err
}

func AddUserLocationAccess(userLocationAccessFormInput *useraccess.User, dbBaa *sql.DB) error {
	// prepare statement to insert values into baa_application.operation.purchase_request
	insertNewUserStr := `
	INSERT INTO  baa_application.operation.pr_location_access (
		fk_user
		,gfk_location)
	VALUES (@p1, @p2)
`
	insertNewUser, err := dbBaa.Prepare(insertNewUserStr)

	res, err := insertNewUser.Exec(
		userLocationAccessFormInput.IDUser,
		userLocationAccessFormInput.GFKLocation,
	)

	log.Println(res)

	return err
}

func GetUserIDEmail(dbBaa *sql.DB) []*useraccess.User {

	stmt, err := dbBaa.Prepare(`
		SELECT 

		pu.id_user
		,pu.email
		
	  	FROM baa_application.operation.pr_user pu
	 `)
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkError(err)
	defer rows.Close()

	userIDEmailTable := []*useraccess.User{}

	for rows.Next() {
		// For each row returned by the table, create a pointer to a User,
		userIDEmail := &useraccess.User{}
		// Populate the attributes of the User,
		// and return incase of an error
		err := rows.Scan(
			&userIDEmail.IDUser,
			&userIDEmail.Email,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		userIDEmailTable = append(userIDEmailTable, userIDEmail)
	}

	return userIDEmailTable

}

func GetDepartmentAccess(dbBaa *sql.DB) []*useraccess.User {

	stmt, err := dbBaa.Prepare(`
		SELECT
  
		ccv.gfk_department
		,ccv.department_name
		
		FROM baa_application.operation.cost_center_view ccv
		GROUP BY ccv.gfk_department, ccv.department_name
	 `)
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkError(err)
	defer rows.Close()

	departmentAccessTable := []*useraccess.User{}

	for rows.Next() {
		// For each row returned by the table, create a pointer to a User,
		departmentAccess := &useraccess.User{}
		// Populate the attributes of the User,
		// and return incase of an error
		err := rows.Scan(
			&departmentAccess.GFKDepartment,
			&departmentAccess.DepartmentName,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		departmentAccessTable = append(departmentAccessTable, departmentAccess)
	}

	return departmentAccessTable

}

func GetLocationAccess(dbBaa *sql.DB) []*useraccess.User {

	stmt, err := dbBaa.Prepare(`
		SELECT
  
		ccv.gfk_location
		,ccv.location_name
		
		FROM baa_application.operation.cost_center_view ccv
		GROUP BY ccv.gfk_location, ccv.location_name
	 `)
	checkError(err)
	defer stmt.Close()

	rows, err := stmt.Query()
	checkError(err)
	defer rows.Close()

	locationAccessTable := []*useraccess.User{}

	for rows.Next() {
		// For each row returned by the table, create a pointer to a User,
		locationAccess := &useraccess.User{}
		// Populate the attributes of the User,
		// and return incase of an error
		err := rows.Scan(
			&locationAccess.GFKDepartment,
			&locationAccess.DepartmentName,
		)
		checkError(err)
		// Finally, append the result to the returned array, and repeat for
		// the next row
		locationAccessTable = append(locationAccessTable, locationAccess)
	}

	return locationAccessTable

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
