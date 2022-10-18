package router

import (
	"final-projek/controllers"
	"final-projek/database"
	"final-projek/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	db := database.StartDB()
	Usercontrollers := &controllers.UserDB{DB: db}
	PhotoControllers := &controllers.PhotoDB{DB: db}
	CommentControllers := &controllers.CommentDB{DB: db}
	SocialMediaControllers := &controllers.SocialMediaDB{DB: db}

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", Usercontrollers.UserRegister)
		userRouter.POST("/login", Usercontrollers.UserLogin)
		userRouter.Use(middlewares.AuthJWT())
		userRouter.PUT("/:userId", Usercontrollers.UserUpdate)
		userRouter.DELETE("/ ", Usercontrollers.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.POST("/", PhotoControllers.CreatePhoto)
		photoRouter.GET("/", PhotoControllers.GetPhotos)
		photoRouter.Use(middlewares.AuthPhoto())
		photoRouter.PUT("/:photoId", PhotoControllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", PhotoControllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.POST("/", CommentControllers.CreateComment)
		commentRouter.GET("/", CommentControllers.GetComments)
		commentRouter.Use(middlewares.AuthComment())
		commentRouter.PUT("/:commentId", CommentControllers.UpdateComment)
		commentRouter.DELETE("/:commentId", CommentControllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.POST("/", SocialMediaControllers.CreateSocialMedia)
		socialMediaRouter.GET("/", SocialMediaControllers.GetSocialMedias)
		socialMediaRouter.Use(middlewares.AuthSocialMedia())
		socialMediaRouter.PUT("/:socialMediaId", SocialMediaControllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", SocialMediaControllers.DeleteSocialMedia)
	}
	return r
}
