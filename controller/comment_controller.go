package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICommentController interface {
	GetAllComments(c echo.Context) error
	GetAllCommentsById(c echo.Context) error
	GetCommentsByArticleID(c echo.Context) error
	CreateComment(c echo.Context) error
	UpdateComment(c echo.Context) error
	DeleteComment(c echo.Context) error
}

type commentController struct {
	cu usecase.ICommentUseCase
}

func NewCommentController(cu usecase.ICommentUseCase) ICommentController {
	return &commentController{cu}
}

func (cc *commentController) GetAllComments(c echo.Context) error {
	comments, err := cc.cu.GetAllComments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, comments)
}

func (cc *commentController) GetAllCommentsById(c echo.Context) error {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	comment, err := cc.cu.GetCommentByID(uint(commentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, comment)
}

func (cc *commentController) GetCommentsByArticleID(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	comments, err := cc.cu.GetCommentsByArticleID(uint(articleId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, comments)
}

func (cc *commentController) CreateComment(c echo.Context) error {
	comment := model.Comment{}
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	commentResponse, err := cc.cu.CreateComment(comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, commentResponse)
}

func (cc *commentController) UpdateComment(c echo.Context) error {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	comment := model.Comment{}
	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	commentResponse, err := cc.cu.UpdateComment(comment, uint(commentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, commentResponse)
}

func (cc *commentController) DeleteComment(c echo.Context) error {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = cc.cu.DeleteComment(uint(commentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Comment deleted")
}
