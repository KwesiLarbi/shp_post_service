package cassandra

import (
    "context"
	"errors"
	client "shp/post_service/internal/database/cassandra"
	"shp/post_service/pkg/models"
)

// Post use from models
type Post models.Post

// Post Service use from models
type PostService models.PostService

// GetPosts get all posts
func (p Post) GetPosts() ([]*models.Post, error) {
    var posts []*models.Post

    ctx := context.Background()
    
    scanner := client.DbClient.Query("SELECT id, abstract, first_name, last_name, historical_interests, post_body, subtitle, title, username, validity_rating, user_id from post_service.posts").WithContext(ctx).Iter().Scanner()

    if scanner != nil {
        return nil, errors.New("no records")
    }

    for scanner.Next() {
        err := scanner.Scan(&p.ID, &p.Username, &p.FirstName, &p.LastName, &p.HistoricalInterests, &p.PostBody, &p.Subtitle, &p.Title, &p.ValidityRating)
        pm := models.Post(p)
        posts = append(posts, &pm)
        if err != nil {
            return nil, err
        } 
    }

    if scanner.Err() != nil {
        return nil, errors.New("no records")
    } 

    return posts, nil
}

// GetPost get one Post
func (p Post) GetPostByID(id string) (*models.Post, error) {
    if err := client.DbClient.Query("SELECT id, abstract, first_name, last_name, historical_interests, post_body, subtitle, title, username, validity_rating, user_id from post_service.posts where id = ?", id).Scan(&p.ID, &p.Username, &p.FirstName, &p.LastName, &p.HistoricalInterests, &p.PostBody, &p.Subtitle, &p.Title, &p.ValidityRating); err != nil {
        return nil, errors.New("no matching records")
    } else {
        post := models.Post(p)
        return &post, nil
    }
}

// CreatePost create a post
func (p *Post) CreatePost() error {
    ctx := context.Background()

    if err := client.DbClient.Query(`INSERT INTO post_service.posts (id, abstract, first_name, last_name, historical_interests, post_body, subtitle, title, username, validity_rating, user_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, p.ID, p.Abstract, p.FirstName, p.LastName, p.HistoricalInterests, p.PostBody, p.Subtitle, p.Title, p.Username, p.ValidityRating, p.UserID).WithContext(ctx).Exec(); err != nil {
        return err
    }
    
    return nil
}