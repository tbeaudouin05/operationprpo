package main

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/operation/operationprpo/api/admin"
	"github.com/thomas-bamilo/operation/operationprpo/api/admin/addcostcenter"
	"github.com/thomas-bamilo/operation/operationprpo/api/admin/adduser"
	"github.com/thomas-bamilo/operation/operationprpo/api/homepage"
	"github.com/thomas-bamilo/operation/operationprpo/api/oauth/authenticate"
	"github.com/thomas-bamilo/operation/operationprpo/api/oauth/login"
	"github.com/thomas-bamilo/operation/operationprpo/api/purchaserequestform"
)

// launch the server and use defined functions to define routes for API calls
func main() {

	router := gin.Default()

	// creating cookie store
	store := sessions.NewCookieStore([]byte(randToken(64)))
	store.Options(sessions.Options{
		Path:   `/`,
		MaxAge: 86400 * 7,
	})

	// ALL THAT IS LEFT IS ADDING THE API CALL TO THIS PAGE!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	// using the cookie store:
	router.Use(sessions.Sessions(`goquestsession`, store))

	router.GET(`/`, homepage.Start)

	router.GET(`/login`, login.LoginHandler)
	router.GET(`/auth`, authenticate.AuthHandler)

	router.GET(`/purchaserequest`, purchaserequestform.Start)
	router.GET(`/purchaserequest/startavailablecostcenter`, purchaserequestform.StartAvailableCostCenter)
	router.GET(`/purchaserequest/startcostcategory`, purchaserequestform.StartCostCategory)
	router.GET(`/purchaserequest/startforminvoicedate`, purchaserequestform.StartInvoiceDate)
	router.POST(`/purchaserequest`, purchaserequestform.AnswerForm)
	router.GET(`/purchaserequest/purchaserequestconfirmation`, purchaserequestform.ConfirmForm)

	router.GET(`/admin`, admin.Start)
	router.GET(`/admin/startformidpurchaserequest`, admin.StartIDPurchaseRequest)
	router.GET(`/admin/startpendingpurchaserequest`, admin.StartPendingPurchaseRequest)
	router.POST(`/admin`, admin.AcceptRejectPurchaseRequest)

	router.GET(`/admin/adduser`, adduser.Start)
	// add user access
	router.GET(`/admin/adduser/startidemail`, adduser.StartIDEmail)
	router.GET(`/admin/adduser/startdepartmentaccess`, adduser.StartDepartmentAccess)
	router.GET(`/admin/adduser/startlocationaccess`, adduser.StartLocationAccess)
	router.POST(`/admin/adduser`, adduser.AnswerForm)
	router.POST(`/adduser/addepartmentaccess`, adduser.AnswerDepartmentAccessForm)
	router.POST(`/adduser/addlocationaccess`, adduser.AnswerLocationAccessForm)
	router.GET(`/admin/adduserconfirmation`, adduser.ConfirmForm)
	// see user access
	router.GET(`/admin/adduser/startExistingUserAccess`, adduser.StartExistingUserAccess)
	router.GET(`/admin/adduser/startExistingUserLocationAccess`, adduser.StartExistingUserLocationAccess)
	router.GET(`/admin/adduser/startExistingUserDepartmentAccess`, adduser.StartExistingUserDepartmentAccess)

	router.GET(`/admin/addcostcenter`, addcostcenter.Start)
	router.GET(`/admin/addcostcenterconfirmation`, addcostcenter.ConfirmForm)
	// cost center
	router.POST(`/addcostcenter/addfunction`, addcostcenter.AnswerCostCenterForm)
	router.GET(`/admin/addcostcenter/startexistingfunctionname`, addcostcenter.StartExistingFunctionName)
	router.GET(`/admin/addcostcenter/startexistingcostcentername`, addcostcenter.StartExistingCostCenterName)
	// department
	router.POST(`/addcostcenter/adddepartment`, addcostcenter.AnswerDepartmentForm)
	router.GET(`/admin/addcostcenter/startexistingdepartmentname`, addcostcenter.StartExistingDepartmentName)
	// location
	router.POST(`/addcostcenter/addlocation`, addcostcenter.AnswerLocationForm)
	router.GET(`/admin/addcostcenter/startexistinglocationname`, addcostcenter.StartExistingLocationName)

	router.Run(`127.0.0.1:8080`)
}

// randToken returns a random token of i bytes
func randToken(i int) string {
	b := make([]byte, i)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
