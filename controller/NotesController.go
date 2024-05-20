package controller

import (
	// "fmt"
	"fmt"
	"gin-temp-app/database"
	"gin-temp-app/types"
	"gin-temp-app/utils"
	"gin-temp-app/views"
	"log"
	"net/http"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AllNotes(c *gin.Context) {
	log.Println("create all note")

	user,err := utils.VerifyToken(c)
	log.Println(user,err,"no user or err")
	if(err != nil || user ==  nil) {
		log.Println("redirecting")
		c.Redirect(http.StatusNotFound, "/login")
	}
	notes := database.Mgr.GetNotes(user.Id, "notes")
    fmt.Println(notes)
	// c.JSON(http.StatusOK, gin.H{
	// 	"error":notes,
	// })
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, views.Base(user.Name,views.AllNotes(notes)))
	c.Render(http.StatusOK, r)
}

func EditNoteForm(c *gin.Context){
	log.Println("create edit form")

	noteId := c.Param("id")
	fmt.Println(noteId)
	_,err := utils.VerifyToken(c)
	if(err != nil) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":"error",
		})
	}
	notes := database.Mgr.GetSingleNoteById(noteId, "notes")
    fmt.Println(notes,"ddd")
	// c.JSON(http.StatusOK, gin.H{
	// 	"error":notes,
	// })
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, views.EditNotes(notes))
	c.Render(http.StatusOK, r)
}

func EditNote(c *gin.Context){
	log.Println("create edit")

	noteId := c.Param("id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	_,err := utils.VerifyToken(c)
	if(err != nil) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":"error",
		})
	}
	data := map[string]interface{}{
		"title": title,
		"content": content,
	}
	erro := database.Mgr.UpdateNote(data, noteId,"notes")
	if erro != nil {
		fmt.Printf("error")
	}
	notes := database.Mgr.GetSingleNoteById(noteId, "notes")
    fmt.Println(notes,"ddd")
	// c.JSON(http.StatusOK, gin.H{
	// 	"error":notes,
	// })
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, views.SingleNote(notes))
	c.Render(http.StatusOK, r)

}

func GetSingleNote(c *gin.Context){
	log.Println("create single")

	noteId := c.Param("id")
	fmt.Println(noteId)
	_,err := utils.VerifyToken(c)
	if(err != nil) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":"error",
		})
	}
	notes := database.Mgr.GetSingleNoteById(noteId, "notes")
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, views.SingleNote(notes))
	c.Render(http.StatusOK, r)

}


func CreateNote(c *gin.Context){
	log.Println("create ,mmmmm")
	user,err := utils.VerifyToken(c)
	log.Println(user,"dlakd;lfal;sd;lfj;alskd")
	if(err != nil){
		log.Println("User not authenticated")
		c.Redirect(http.StatusUnauthorized, "/login")
	}
	r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, views.Base(user.Name, views.CreateNoteForm()) )
	c.Render(http.StatusOK, r)

}


func SaveNote(c *gin.Context){
	log.Println("create save ntoe")

	user,err := utils.VerifyToken(c)
	title, content :=c.PostForm("title"),c.PostForm("content")
	if(err != nil){
		log.Println("User not authenticated")
		c.Redirect(http.StatusUnauthorized, "/login")
	}
	id,_ := primitive.ObjectIDFromHex(user.Id.Hex())
	
	// data := 
	resultId, err := database.Mgr.Insert(types.Note{
		Title:title,
		Content: content,
		UserID: id ,
	},"notes")

	if(err != nil){
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Could not create note",
		})
	}
    log.Println(resultId)
	c.Redirect(http.StatusFound, "/notes", )


}