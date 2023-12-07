package repository

import (
	"context"

	"github.com/fidya02/Capstone-Project/entity"

	"gorm.io/gorm"
)

// Blog repository
type BlogRepository struct {
	db *gorm.DB
}

// NewBlogRepository creates a new instance of BlogRepository.
// It takes a *gorm.DB as a parameter and returns a pointer to BlogRepository.
func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

// GetAllBlogs retrieves all blogs from the database.
// It returns a slice of *entity.Blog and an error.
func (r *BlogRepository) GetAllBlogs(ctx context.Context) ([]*entity.Blog, error) {
	// Initialize an empty slice to store the blogs
	blogs := make([]*entity.Blog, 0)

	// Execute the query to retrieve all blogs from the database
	result := r.db.WithContext(ctx).Find(&blogs)

	// Check if there was an error executing the query
	if result.Error != nil {
		return nil, result.Error
	}

	// Return the retrieved blogs and nil error
	return blogs, nil
}

// CreateBlog saves a new Blog to the database.
func (r *BlogRepository) CreateBlog(ctx context.Context, blog *entity.Blog) error {
	// Use the context to create the blog record in the database
	result := r.db.WithContext(ctx).Create(blog)

	// Check if there was an error while creating the blog record
	if result.Error != nil {
		// Return the error if there was one
		return result.Error
	}

	// Return nil if there were no errors
	return nil
}

// UpdateBlog updates a Blog in the database.
func (r *BlogRepository) UpdateBlog(ctx context.Context, blog *entity.Blog) error {
	// Update the blog in the database
	result := r.db.WithContext(ctx).Model(&entity.Blog{}).Where("id = ?", blog.ID).Updates(&blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetBlog retrieves a Blog by its ID from the database.
func (r *BlogRepository) GetBlog(ctx context.Context, id int64) (*entity.Blog, error) {
	Blog := new(entity.Blog)
	result := r.db.WithContext(ctx).First(&Blog, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return Blog, nil
}

// DeleteBlog deletes a Blog from the database.
func (r *BlogRepository) DeleteBlog(ctx context.Context, id int64) error {
	result := r.db.WithContext(ctx).Delete(&entity.Blog{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SearchBlog search Blog
func (r *BlogRepository) SearchBlog(ctx context.Context, search string) ([]*entity.Blog, error) {
	blogs := make([]*entity.Blog, 0)
	result := r.db.WithContext(ctx).Where("title LIKE ?", "%"+search+"%").Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}
