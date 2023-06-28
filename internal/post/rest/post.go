package rest

import (
	"fmt"
	"net/http"
	db "shp/post_service/internal/post/repo/cassandra"
	"shp/post_service/internal/post/service"

	"github.com/labstack/echo/v4"
)

// GetPosts responds with the list of all posts as JSON.
func GetPosts(c echo.Context) error {
	response, err := service.GetPosts()

	if err != nil {
		panic(err)
	}
	
    return c.JSON(http.StatusOK, response)
}

// GetPostByID locates the post whose ID value matches the id
// parameter sent by the client, then returns that post as a response.
func GetPostByID(c echo.Context) error {
    id := c.Param("id")

	response, err:= service.GetPostByID(id)

	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, response)
}

// CreatePost adds a post from JSON received in the request body.
func CreatePost(c echo.Context) error {
	fmt.Println("Creating a Post")

    p := new(db.Post)

    // Bind request body to the post model object
    if err := c.Bind(p); err != nil {
        panic(err)
    }

	err := service.CreatePost(p)

	if err != nil {
		panic(err)
	}
    
    return c.String(http.StatusCreated, "Post created successfully")
}