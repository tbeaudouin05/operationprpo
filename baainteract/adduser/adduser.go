package adduser

import (
	"database/sql"
	"log"

	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"
)

// LoadPurchaseRequestToDb loads purchase request to baa_application.operation.purchase_request
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

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
