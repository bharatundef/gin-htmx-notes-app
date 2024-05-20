package router

import (
	// "fmt"
	// "gin-temp-app/controller"
	// "net/http"
	// "os"

	"gin-temp-app/controller"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name       string
	Method     string
	Path       string // localhost:8000/helloworld
	HandleFunc func(*gin.Context)
}

type Routes []Route


type routes struct {
	router *gin.Engine

} 

func Routing(){
	r := routes{
		router: gin.Default(),
	}
	r.router.SetTrustedProxies(nil)
	// grouping := r.router.Group(os.Getenv("ApiVersion"))
	// grouping := r.router.Group("/")
	// r.HandleRoutes(grouping)
	r.router.GET("/", controller.HelloUser )
	r.router.GET("/login", controller.LoginUserForm )
	r.router.GET("/logout", controller.LogoutUser )
	r.router.POST("/user/login", controller.LoginUser )
	r.router.POST("/user/register", controller.RegisterUser )
	r.router.GET("/register", controller.RegisterUserForm )
	r.router.GET("/notes", controller.AllNotes )

	r.router.GET("/notes/create", controller.CreateNote)
	r.router.POST("/notes/save", controller.SaveNote )
	r.router.GET("/notes/:id", controller.GetSingleNote )
	r.router.GET("/notes/:id/edit", controller.EditNoteForm )
	r.router.PUT("/notes/:id/edit", controller.EditNote )
	r.router.Run(":" + os.Getenv("Port"))
}