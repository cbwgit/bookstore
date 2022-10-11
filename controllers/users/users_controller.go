package users

import (
	"bookstore/domain/users"
	"bookstore/services"
	
	"net/http"

	"bookstore/utils/errors"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

func Test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invaild json body")
		
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
		//ToDo: handle user creation error
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
		
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
