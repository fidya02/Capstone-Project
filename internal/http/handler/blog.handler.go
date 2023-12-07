package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fidya02/Capstone-Project/entity"
	"github.com/fidya02/Capstone-Project/internal/http/validator"
	"github.com/fidya02/Capstone-Project/internal/service"
	"github.com/labstack/echo/v4"
)

// BlogHandler handles HTTP requests related to Blogs.
type BlogHandler struct {
	blogService service.BlogUseCase
}

// NewBlogHandler creates a new instance of BlogHandler.
// It takes a blogService as a parameter and returns a pointer to a BlogHandler.
func NewBlogHandler(blogService service.BlogUseCase) *BlogHandler {
	// Create a new instance of BlogHandler, initialize its blogService field with the provided blogService, and return the pointer to the new instance.
	return &BlogHandler{blogService}
}

// GetAllBlogs is a handler function that retrieves all blogs.
// It takes in an echo.Context object and returns an error or a JSON response.
func (h *BlogHandler) GetAllBlogs(c echo.Context) error {
	// Call the GetAllBlogs method of the blogService to retrieve all blogs
	blogs, err := h.blogService.GetAllBlogs(c.Request().Context())
	if err != nil {
		// Return an error response with status code 422 if an error occurred
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	// Return a JSON response with status code 200 and the retrieved blogs
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": blogs,
	})
}

// CreateBlog creates a new blog using the provided data in the request body.
func (h *BlogHandler) CreateBlog(c echo.Context) error {
	// Define the input structure to hold the data from the request body
	var input struct {
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Date        time.Time `json:"date"`
		Image       string    `json:"image"`
	}

	// Input validation
	if err := c.Bind(&input); err != nil {
		// Return a bad request error with the validation errors
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Convert input.Date to a string with the desired format
	dateString := input.Date.Format("2006-01-02T15:04:05Z")

	// Create a Blog object using the input data
	blog := entity.Blog{
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Date:        dateString,
	}

	// Call the blogService to create the Blog
	err := h.blogService.CreateBlog(c.Request().Context(), &blog)
	if err != nil {
		// Return an unprocessable entity error if there was an error creating the Blog
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// Return a success message
	return c.JSON(http.StatusCreated, "Blog created successfully")
}

// GetBlog handles the retrieval of a Blog by ID.
func (h *BlogHandler) GetBlog(c echo.Context) error {
	idStr := c.Param("id")                     // assuming the ID is passed as a URL parameter as a string
	id, err := strconv.ParseInt(idStr, 10, 64) // Convert the string to int64
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid ID",
		})
	}

	Blog, err := h.blogService.GetBlog(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"id":          Blog.ID,
			"title":       Blog.Title,
			"description": Blog.Description,
			"image":       Blog.Image,
			"date":        Blog.Date,
			"created":     Blog.CreatedAt,
		},
	})
}

// UpdateBlog handles the update of an existing Blog.
func (h *BlogHandler) UpdateBlog(c echo.Context) error {
	var input struct {
		ID          int64     `param:"id" validate:"required"`
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		Image       string    `json:"image"`
		Date        time.Time `json:"date"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	// Convert input.Date to a formatted string
	dateStr := input.Date.Format("2006-01-02T15:04:05Z")

	// Create a Blog object
	Blog := entity.Blog{
		ID:          input.ID, // Assuming ID is already of type int64
		Title:       input.Title,
		Description: input.Description,
		Image:       input.Image,
		Date:        dateStr, // Assign the formatted date string
	}

	err := h.blogService.UpdateBlog(c.Request().Context(), &Blog)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Blog updated successfully",
		"Blog":    Blog,
	})
}

// DeleteBlog handles the deletion of a Blog by ID.
func (h *BlogHandler) DeleteBlog(c echo.Context) error {
	var input struct {
		ID int64 `param:"id" validate:"required"`
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	err := h.blogService.DeleteBlog(c.Request().Context(), input.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Blog deleted successfully",
	})
}

// SearchBlog handles the search of a Blog by title.
func (h *BlogHandler) SearchBlog(c echo.Context) error {
	var input struct {
		Search string `param:"search" validate:"required"` //harus pramater search
	}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	Blogs, err := h.blogService.SearchBlog(c.Request().Context(), input.Search)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": Blogs,
	})
}
