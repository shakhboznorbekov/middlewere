package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "crud/api/docs"
	"crud/api/handler"
	"crud/storage"
)

func SetUpApi(r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandlerV1(storage)

	r.Use(customCORSMiddleware())

	v1 := r.Group("/v1")

	v1.Use(checkPassword())
	v1.GET("/order", handlerV1.GetOrderList)

	r.POST("/order", handlerV1.CreateOrder)
	r.GET("/order/:id", handlerV1.GetOrderById)
	r.PUT("/order/:id", handlerV1.UpdateOrder)
	r.DELETE("/order/:id", handlerV1.DeleteOrder)

	r.POST("/user", handlerV1.CreateUser)
	r.GET("/user/:id", handlerV1.GetUserById)
	r.GET("/user", handlerV1.GetUserList)
	r.PUT("/user/:id", handlerV1.UpdateUser)
	r.DELETE("/user/:id", handlerV1.DeleteUser)

	r.POST("/book", handlerV1.CreateBook)
	r.GET("/book/:id", handlerV1.GetBookById)
	r.GET("/book", handlerV1.GetBookList)
	r.PUT("/book/:id", handlerV1.UpdateBook)
	r.DELETE("/book/:id", handlerV1.DeleteBook)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func checkPassword() gin.HandlerFunc {

	return func(c *gin.Context) {

		if _, ok := c.Request.Header["Password"]; ok {
			if c.Request.Header["Password"][0] != "1234" {
				c.AbortWithError(http.StatusForbidden, errors.New("not found password"))
				return
			}
		} else {
			c.AbortWithError(http.StatusForbidden, errors.New("not found password"))
			return
		}

		c.Next()
	}
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
