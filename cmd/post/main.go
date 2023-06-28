package main

import (
	"shp/post_service/internal/post/rest"
    
	"github.com/labstack/echo/v4"
)

var (
	router = echo.New()
)

func main() {
    mapUrls()
    router.Logger.Fatal(router.Start("8080"))
}

func mapUrls() {
    router.GET("/posts", rest.GetPosts)
    router.GET("/posts/:id", rest.GetPostByID)
    router.GET("/posts/create_post", rest.CreatePost)
}