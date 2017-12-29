package router

import (
	"mission-board/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")

		if token == "" {
			c.JSON(401, gin.H{
				"message": "Token required",
			})
			c.Abort()
			return
		}

		if token != "accesstoken" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Invalid Token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func Init() {
	// Creates a default gin router
	// PrintMiddleware is a function for test middleware

	r := gin.Default() // Grouping routes
	// group： v1
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", handlers.HelloPage)
		v1.GET("/hello/:name/:age", func(c *gin.Context) {
			name := c.Param("name")
			age := c.Param("age")
			c.String(http.StatusOK, "Hello %s ,age is %s", name, age)
		})
		v1.GET("/welcome", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "guest")
			lastname := c.Query("lastname")
			c.String(http.StatusOK, "welcome %s , %s", firstname, lastname)
		})
		v1.GET("abTest",handlers.TestAb)
		v1.GET("abTest2",handlers.TestAb2)
	}


	// 404 NotFound
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	v1.GET("/emergencyNotice",handlers.EmergencyNotice)


	r.Run(":8090") // listen and serve on 0.0.0.0:8000
}
