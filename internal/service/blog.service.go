package service

import (
	"context"

	"github.com/fidya02/Capstone-Project/entity"
)

// BlogUseCase is an interface for Blog-related use cases.
type BlogUseCase interface {
	GetAllBlogs(ctx context.Context) ([]*entity.Blog, error)
	CreateBlog(ctx context.Context, Blog *entity.Blog) error
	GetBlog(ctx context.Context, id int64) (*entity.Blog, error)
	UpdateBlog(ctx context.Context, Blog *entity.Blog) error
	SearchBlog(ctx context.Context, search string) ([]*entity.Blog, error)
	DeleteBlog(ctx context.Context, id int64) error
}

type BlogRepository interface {
	GetAllBlogs(ctx context.Context) ([]*entity.Blog, error)
	CreateBlog(ctx context.Context, Blog *entity.Blog) error
	GetBlog(ctx context.Context, id int64) (*entity.Blog, error)
	UpdateBlog(ctx context.Context, Blog *entity.Blog) error
	SearchBlog(ctx context.Context, search string) ([]*entity.Blog, error)
	DeleteBlog(ctx context.Context, id int64) error
}

// BlogService is responsible for Blog-related business logic.
type BlogService struct {
	Repository BlogRepository
}

// NewBlogService creates a new instance of BlogService.
// It takes a BlogRepository as a parameter and returns a pointer to a BlogService.
func NewBlogService(Repository BlogRepository) *BlogService {
	return &BlogService{Repository: Repository}
}

// GetAllBlogs retrieves all blogs from the repository.
func (s *BlogService) GetAllBlogs(ctx context.Context) ([]*entity.Blog, error) {
	// Call the GetAllBlogs method of the repository to get the blogs.
	blogs, err := s.Repository.GetAllBlogs(ctx)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

// CreateBlog creates a new blog.
func (s *BlogService) CreateBlog(ctx context.Context, blog *entity.Blog) error {
	return s.Repository.CreateBlog(ctx, blog)
}

// UpdateBlog updates a blog in the database.
func (s *BlogService) UpdateBlog(ctx context.Context, blog *entity.Blog) error {
	return s.Repository.UpdateBlog(ctx, blog)
}

// GetBlog retrieves a blog with the specified ID from the repository.
// It returns the blog entity and an error if any.
func (s *BlogService) GetBlog(ctx context.Context, id int64) (*entity.Blog, error) {
	return s.Repository.GetBlog(ctx, id)
}

// DeleteBlog deletes a blog with the given ID.
// It returns an error if the deletion fails.
func (s *BlogService) DeleteBlog(ctx context.Context, id int64) error {
	return s.Repository.DeleteBlog(ctx, id)
}

// SearchBlog searches for blogs based on a search string.
//
// It takes a context and a search string as input.
// It returns a slice of blog entities and an error, if any.
func (s *BlogService) SearchBlog(ctx context.Context, search string) ([]*entity.Blog, error) {
	return s.Repository.SearchBlog(ctx, search)
}
