package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	//Public Routes
	public := r.Group("/")
	public.POST("/api/login", controllers.Login)
	public.POST("/averifypi/register", controllers.Register)
	public.GET("/api/logout", controllers.Logout)

	//=============================Middlewares for Siswa======================================================
	siswa := r.Group("/api/siswa")
	//nanti ini kita isi middlewarenya
	siswa.GET("/user", controllers.GetUser)
	siswa.GET("/user/:id", controllers.GetUserById)

	return r
}
