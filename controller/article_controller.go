package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IArticleInterface interface {
	GetAllArticles(c echo.Context) error
	GetAllArticlesById(c echo.Context) error
	CreateArticle(c echo.Context) error
	UpdateArticle(c echo.Context) error
	DeleteArticle(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type articleController struct {
	au usecase.IArticleUseCase
}

func NewArticleController(au usecase.IArticleUseCase) IArticleInterface {
	return &articleController{au}
}

func (ac *articleController) GetAllArticles(c echo.Context) error {
	articles, err := ac.au.GetAllArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, articles)
}

func (ac *articleController) GetAllArticlesById(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	article, err := ac.au.GetArticleByID(uint(articleId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, article)
}

func (ac *articleController) CreateArticle(c echo.Context) error {
	article := model.Article{}
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	articleResponse, err := ac.au.CreateArticle(article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, articleResponse)
}

func (ac *articleController) UpdateArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	article := model.Article{}
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	articleResponse, err := ac.au.UpdateArticle(article, uint(articleId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, articleResponse)
}

func (ac *articleController) DeleteArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = ac.au.DeleteArticle(uint(articleId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Article deleted")
}

func (ac *articleController) CsrfToken(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"csrfToken": c.Get("csrf").(string),
	})
}
