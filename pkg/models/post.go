package models

type Post struct {
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

// Post Service interface for Post Model
type PostService interface {
	GetPosts() ([]*Post, error)
	GetPostByID(id string) (*Post, error)
	CreatePost() error
}