package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Historical interest struct for posts
// TODO: Define interest string in database
// TODO: Show interest relationship in ER diagram
// type historicalInterests struct {
//     name string
// }

// post represesnts data about a singular post from a user
type post struct {
    ID                      string      `json:"id"`
    Username                string      `json:"username"`
    FirstName               string      `json:"first_name"`
    LastName                string      `json:"last_name"`
    Title                   string      `json:"title"`
    Subtitle                string      `json:"subtitle"`
    Abstract                string      `json:"abstract"`
    PostBody                string      `json:"post_body"`
    ValidityRating          float64     `json:"validity_rating"`
    HistoricalInterests     []string    `json:"historical_interests"`
    UserID                  string      `json:"user_id"`
}

// posts slice to seed post data
var posts = []post{
    {
        ID: "1",
        Username: "jdoe",
        FirstName: "John",
        LastName: "Doe",
        Title: "First Seed Post",
        Subtitle: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
        Abstract: "Massa sapien faucibus et molestie. Tempor commodo ullamcorper a lacus vestibulum sed. Tristique senectus et netus et malesuada fames. Feugiat nisl pretium fusce id.",
        PostBody: "Euismod quis viverra nibh cras pulvinar. Viverra vitae congue eu consequat ac felis. Diam phasellus vestibulum lorem sed risus ultricies tristique nulla aliquet. Quisque non tellus orci ac auctor augue mauris. Elit ullamcorper dignissim cras tincidunt lobortis feugiat vivamus at. Cursus euismod quis viverra nibh. Ullamcorper velit sed ullamcorper morbi tincidunt ornare massa. Aliquam ut porttitor leo a. Vulputate odio ut enim blandit volutpat. Consectetur libero id faucibus nisl tincidunt eget nullam non. Aliquam eleifend mi in nulla posuere sollicitudin. Morbi tempus iaculis urna id. Vel pretium lectus quam id leo.",
        ValidityRating: 100.00,
        HistoricalInterests: []string{"truth", "test", "history"},
        UserID: "1",
    },
    {
        ID: "2",
        Username: "janeTheDame",
        FirstName: "Jane",
        LastName: "Doe",
        Title: "Second Seed Post",
        Subtitle: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
        Abstract: "Massa sapien faucibus et molestie. Tempor commodo ullamcorper a lacus vestibulum sed. Tristique senectus et netus et malesuada fames. Feugiat nisl pretium fusce id.",
        PostBody: "Euismod quis viverra nibh cras pulvinar. Viverra vitae congue eu consequat ac felis. Diam phasellus vestibulum lorem sed risus ultricies tristique nulla aliquet. Quisque non tellus orci ac auctor augue mauris. Elit ullamcorper dignissim cras tincidunt lobortis feugiat vivamus at. Cursus euismod quis viverra nibh. Ullamcorper velit sed ullamcorper morbi tincidunt ornare massa. Aliquam ut porttitor leo a. Vulputate odio ut enim blandit volutpat. Consectetur libero id faucibus nisl tincidunt eget nullam non. Aliquam eleifend mi in nulla posuere sollicitudin. Morbi tempus iaculis urna id. Vel pretium lectus quam id leo.",
        ValidityRating: 100.00,
        HistoricalInterests: []string{"truth", "test", "history"},
        UserID: "2",
    },
    {
        ID: "3",
        Username: "klarbi",
        FirstName: "Kwesi",
        LastName: "Larbi",
        Title: "Third Seed Post",
        Subtitle: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
        Abstract: "Massa sapien faucibus et molestie. Tempor commodo ullamcorper a lacus vestibulum sed. Tristique senectus et netus et malesuada fames. Feugiat nisl pretium fusce id.",
        PostBody: "Euismod quis viverra nibh cras pulvinar. Viverra vitae congue eu consequat ac felis. Diam phasellus vestibulum lorem sed risus ultricies tristique nulla aliquet. Quisque non tellus orci ac auctor augue mauris. Elit ullamcorper dignissim cras tincidunt lobortis feugiat vivamus at. Cursus euismod quis viverra nibh. Ullamcorper velit sed ullamcorper morbi tincidunt ornare massa. Aliquam ut porttitor leo a. Vulputate odio ut enim blandit volutpat. Consectetur libero id faucibus nisl tincidunt eget nullam non. Aliquam eleifend mi in nulla posuere sollicitudin. Morbi tempus iaculis urna id. Vel pretium lectus quam id leo.",
        ValidityRating: 100.00,
        HistoricalInterests: []string{"truth", "test", "history"},
        UserID: "3",
    },
}

func main() {
    e := echo.New()
    e.GET("/posts", getPosts)
    e.GET("/posts/:id", getPostByID)
    e.GET("/posts/create_post", createPost)

    e.Logger.Fatal(e.Start(":8080"))
}

// getPosts responds with the list of all posts as JSON.
func getPosts(c echo.Context) error {
    return c.JSON(http.StatusOK, posts)
}

// getPostByID locates the post whose ID value matches the id
// parameter sent by the client, then returns that post as a response.
func getPostByID(c echo.Context) error {
    id := c.Param("id")

    // Loop through thet list of posts, looking for an post whose ID value
    // matches the parameter.
    for _, a := range posts {
        if a.ID == id {
            return c.JSON(http.StatusOK, a)
        }
    }

    return c.String(http.StatusOK, id)
}

// createPost adds a post from JSON received in the request body.
func createPost(c echo.Context) error {
    var newPost post

    // Call BindJSON to bind the received JSON to post
    if err := c.Bind(&newPost); err != nil {
        return err
    }

    // add the new post to the slice
    posts = append(posts, newPost)
    
    return c.JSON(http.StatusCreated, newPost)
}