package main

import (
	blogDetails "blog-app-service/internal/api/blog_details"
	blogPosts "blog-app-service/internal/api/blog_posts"
	"blog-app-service/internal/api/login"
	"blog-app-service/internal/dal"
	"blog-app-service/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	DB, err := database.Init()
	if err != nil {
		panic(err)
	}

	login := &login.Handler{
		LoginDAO: dal.NewLoginDAO(DB),
	}
	blogPosts := &blogPosts.Handler{
		BlogPostsDAO: dal.NewBlogPostsDAO(DB),
	}
	blogDetails := &blogDetails.Handler{
		BlogDetailsDAO: dal.NewBlogDetailsDAO(DB),
	}

	router := gin.Default()
	api := router.Group("/api")
	api.POST("/login", login.LoginHandler)
	api.GET("/blogPosts", blogPosts.BlogPostsHandler)
	api.GET("/blogPost/:id", blogDetails.BlogDetailsHandler)

	router.Run("localhost:8080")
}
