package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func CreateUser(c *gin.Context) {
// 	var user users.User

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		//TODO: return bad request to the caller.
// 		restErr := errors.NewBadRequestError("invalid JSON body")
// 		c.JSON(restErr.Status, restErr)
// 		return
// 	}
// 	result, saveErr := services.CreateUser(user)

// 	if saveErr != nil {
// 		//TODO: handle user creation error
// 		c.JSON(saveErr.Status, saveErr)
// 		return
// 	}

// 	c.JSON(http.StatusCreated, result)
// }

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented")
}
