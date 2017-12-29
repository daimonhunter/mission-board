package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
func PostLogin(c *gin.Context) {
	var form Login
	// 根据请求头中 content-type 自动推断.
	if c.Bind(&form) == nil {
		if form.User == "manc" && form.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
}