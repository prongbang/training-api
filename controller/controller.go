package controller

import (
	"net/http"
	"strconv"
	"time"
	"training-api/model"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Controller struct {
	DB *gorm.DB
}

func (ctrl Controller) GetHome(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{"message": "Hello, World!"})
}

func (ctrl Controller) GetPostsAll(c echo.Context) error {

	var posts []model.Post
	ctrl.DB.Find(&posts)

	return c.JSON(http.StatusOK, posts)
}

func (ctrl Controller) GetPostById(c echo.Context) error {

	idstr := c.Param("id")
	if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {
		var posts model.Post
		ctrl.DB.Where("id = ?", id).Find(&posts)
		if posts.ID > 0 {
			return c.JSON(http.StatusOK, posts)
		}
		return c.JSON(http.StatusOK, echo.Map{})
	}
	return c.JSON(http.StatusNotFound, echo.Map{"message": "Not found"})
}

func (ctrl Controller) PostPost(c echo.Context) error {

	var post model.Post
	if c.Bind(&post) == nil {
		var last model.Post
		ctrl.DB.Last(&last)
		var id int64 = 1
		if last.ID > 0 {
			id = last.ID + 1
		}
		post.ID = id
		post.CreatedAt = time.Now()
		ctrl.DB.Create(&post)
		return c.JSON(http.StatusCreated, post)
	}
	return c.JSON(http.StatusBadRequest, echo.Map{"message": "Bad request"})
}

func (ctrl Controller) PutPost(c echo.Context) error {
	idstr := c.Param("id")
	var post model.Post
	if c.Bind(&post) == nil {
		if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {
			var p model.Post
			ctrl.DB.Where("id = ?", id).First(&p)
			if p.ID > 0 {
				// Update
				p.UpdatedAt = time.Now()
				p.Subject = post.Subject
				p.Content = post.Content
				ctrl.DB.Model(&post).Where("id = ?", id).Update(&p)
			} else {
				// Create
				p.ID = id
				p.CreatedAt = time.Now()
				p.Subject = post.Subject
				p.Content = post.Content
				ctrl.DB.Create(&p)
			}
			return c.JSON(http.StatusOK, p)
		}
	}
	return c.JSON(http.StatusBadRequest, echo.Map{"message": "Bad request"})
}

func (ctrl Controller) PatchPost(c echo.Context) error {
	idstr := c.Param("id")
	var post model.Post
	if c.Bind(&post) == nil {
		if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {
			var p model.Post
			ctrl.DB.Where("id = ?", id).First(&p)
			if p.ID > 0 {
				// Update
				p.UpdatedAt = time.Now()
				if post.Subject != "" {
					p.Subject = post.Subject
				}
				if post.Content != "" {
					p.Content = post.Content
				}
				ctrl.DB.Model(&post).Where("id = ?", id).Update(&p)
				return c.JSON(http.StatusOK, p)
			}
			return c.JSON(http.StatusNotFound, echo.Map{"message": "Not found"})
		}
	}
	return c.JSON(http.StatusBadRequest, echo.Map{"message": "Bad request"})
}

func (ctrl Controller) DeletePost(c echo.Context) error {
	idstr := c.Param("id")
	if id, err := strconv.ParseInt(idstr, 10, 64); err == nil {
		var p model.Post
		ctrl.DB.Where("id = ?", id).First(&p)
		if p.ID > 0 {
			ctrl.DB.Where("id = ?", id).Delete(&p)
			return c.JSON(http.StatusOK, p)
		}
		return c.JSON(http.StatusNotFound, echo.Map{"message": "Not found"})
	}
	return c.JSON(http.StatusBadRequest, echo.Map{"message": "Bad request"})
}
