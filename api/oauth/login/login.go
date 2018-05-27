package login

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/thomas-bamilo/operation/operationprpo/row/useraccess"

	"github.com/thomas-bamilo/operation/operationprpo/api/oauth/credential"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var user useraccess.User
	state := randToken(32)
	session := sessions.Default(c)
	session.Set("state", state)
	session.Save()
	user.LoginURL = getLoginURL(state)
	user.Render(c, `template/login/login.html`)
}

func getLoginURL(state string) string {
	conf := credential.InitCred()
	return conf.AuthCodeURL(state)
}

// randToken returns a random token of i bytes
func randToken(i int) string {
	b := make([]byte, i)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
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
