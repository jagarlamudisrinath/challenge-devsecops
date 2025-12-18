package api

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {

	r := gin.Default()

	// route v1
	v1 := r.Group("/v1")
	{
		v1.GET("users", ListUsers)
		v1.POST("user/add", AddUser)
		v1.GET("user/:id", GetUser)
		v1.GET("user/delete/:id", DelUser)
	}

	return r

}
