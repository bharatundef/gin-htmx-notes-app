package controller

import (
	// "fmt"
	"net/http"

	// "gin-temp-app/database"
	// "gin-temp-app/models"
	"gin-temp-app/database"
	"gin-temp-app/types"
	"gin-temp-app/utils"
	"gin-temp-app/views"
	"gin-temp-app/views/userT"
	"log"

	"fmt"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func HelloUser(c *gin.Context) {
	// ginHtmlRenderer := router.GinDefault.rendere
	// user := models.User{}
	// err := c.BindJSON(&user)
	// if err != nil {
	// 	fmt.Println("Can't bind data")
	// 	return
	// }
    // fmt.Println(user)
	// database
	// err = database.InsertData(user)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": "error occured"})
	// 	return
	// }
	// c.HTML(http.StatusOK, templ.Handler(views.))
	
	// c.HTML(http.StatusOK, "", userT.UserP())

	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, views.Base("",userT.UserP("name")))
	c.Render(http.StatusOK, r)
	// c.Render(http.StatusOK, r)
}


func RegisterUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	name := c.PostForm("name")
	fmt.Println(email,password)
	alreadyUSER := database.Mgr.GetSingleRecordByUsername(email, "user")
	if(alreadyUSER.Email != "" ){
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})

	}
	loggedInUser, err := database.Mgr.Insert(types.User{
		Name: name,
		Email: email,
		Password: password,
	}, "user")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid login data"})
        
	}
    log.Println(loggedInUser,"loggenintu")
	
	utils.CreateToken(c, email)
	

	// http.SetCookie(c.Writer, &http.Cookie{
    //     Name:     "auth_token",
    //     Value:    signedToken,
    //     Expires:  time.Now().Add(time.Hour * 24 * 7), 
    //     HttpOnly: false, 
    //     Secure:   false, 
    // })
    log.Println("Created token")
    c.Redirect(http.StatusFound, "/notes")

}

func LoginUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	
	loggedInUser := database.Mgr.GetUserByEmailPassword(email,password, "user")
	if loggedInUser.Email != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
	}
    log.Println(loggedInUser,"loggenintu")
	
	utils.CreateToken(c, email)


    c.Redirect(http.StatusFound, "/notes")

}


func LoginUserForm(c *gin.Context) {
	// cookie,err := c.Cookie("auth_token")
	// if(err != nil || cookie != ""){
	// 	_,err := utils.VerifyToken(c)
	// 	if(err != nil){
	// 		c.Redirect(http.StatusMovedPermanently, "/notes")     
	// 	}
	// }
    
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, views.Base("",views.Login()))
	c.Render(http.StatusOK, r)

}


func RegisterUserForm(c *gin.Context) {
	cookie,err := c.Cookie("auth_token")
    log.Println("inside register user form")
    log.Println(len(cookie),err,"err")
	if(cookie == ""){
		log.Println("error detected")
		r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, views.Base("",views.RegisterUserForm()))
		c.Render(http.StatusOK, r)	
		// c.Abort()
		return
	}
	log.Println("error after")

	_,err2 := utils.VerifyToken(c)

	log.Println("error verify")

	if(err2 != nil){
		log.Println("error note nil")

		c.Redirect(http.StatusUnauthorized, "/")   
		return  
	}


	// if(err2 != nil){
	log.Println("error redirect")
	
	c.Redirect(http.StatusMovedPermanently, "/notes")     
	// }
    
	

}

func LogoutUser(c *gin.Context){
	log.Println("logutuser")
	c.SetCookie("auth_token", "", -1, "/", "localhost", false, false)
	c.Redirect(http.StatusFound, "/login")
	c.Abort()
}