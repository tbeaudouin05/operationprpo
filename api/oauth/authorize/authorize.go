package authorize

import (
	"fmt"
	"net/http"

	"github.com/thomas-bamilo/operation/operationprpo/baainteract"
	"github.com/thomas-bamilo/sql/connectdb"

	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthorizeCcAdmin(c *gin.Context, user *useraccess.User) {

	session := sessions.Default(c)
	userEmail := session.Get("userEmail")
	if userEmail == nil {
		http.Redirect(c.Writer, c.Request, `/login`, http.StatusSeeOther)
		return
	}

	user.Email = userEmail.(string)

	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()
	baainteract.GetUserInfo(user, dbBaa)
	if user.Access != `cc_admin` && user.Access != `admin` {
		session.Set("unauthorized", "You need an admin account to access this web page!")
		err := session.Save()
		handleErr(c, err)
		http.Redirect(c.Writer, c.Request, `/`, http.StatusSeeOther)
		return
	}

	session.Set("userName", user.Name)

}

func AuthorizePrAdmin(c *gin.Context, user *useraccess.User) {

	session := sessions.Default(c)
	userEmail := session.Get("userEmail")
	if userEmail == nil {
		http.Redirect(c.Writer, c.Request, `/login`, http.StatusSeeOther)
		return
	}

	user.Email = userEmail.(string)

	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()
	baainteract.GetUserInfo(user, dbBaa)
	if user.Access != `pr_admin` && user.Access != `cc_admin` && user.Access != `admin` {
		session.Set("unauthorized", "You need an admin account to access this web page!")
		err := session.Save()
		handleErr(c, err)
		http.Redirect(c.Writer, c.Request, `/`, http.StatusSeeOther)
		return
	}

	session.Set("userName", user.Name)

}

func Authorize(c *gin.Context, user *useraccess.User) {

	session := sessions.Default(c)
	userEmail := session.Get("userEmail")
	if userEmail == nil {
		http.Redirect(c.Writer, c.Request, `/login`, http.StatusSeeOther)
		return
	}

	user.Email = userEmail.(string)

	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()
	baainteract.GetUserInfo(user, dbBaa)

	if user.Access == `` {
		session.Set("unauthorized", "Your account cannot access this application!")
		err := session.Save()
		handleErr(c, err)
		http.Redirect(c.Writer, c.Request, `/unauthorized`, http.StatusSeeOther)
		return
	}

	session.Set("userName", user.Name)

}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
