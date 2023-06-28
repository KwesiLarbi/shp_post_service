package service

import (
	db "shp/post_service/internal/post/repo/cassandra"
	"shp/post_service/pkg/models"
)

var ps db.PostService
var post db.Post

// GetPosts get Post
func GetPosts() ([]*models.Post, error) {
	ps = &post
    posts, err := ps.GetPosts()
    if err != nil {
        return nil, err
    }
    return posts, nil
}

// GetPostByID get one Post
func GetPostByID(id string) (*models.Post, error) {
    ps = &post
    posts, err := ps.GetPostByID(id)
    if err != nil {
        return nil, err
    }
    return posts, nil
}

// CreatePost creates a Post
func CreatePost(p *db.Post) error {
    ps = p
    err := ps.CreatePost()
    if err != nil {
        return err
    }
    return nil
}