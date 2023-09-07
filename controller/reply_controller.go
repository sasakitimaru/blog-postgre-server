package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IReplyController interface {
	GetAllReplies(c echo.Context) error
	GetAllRepliesById(c echo.Context) error
	GetRepliesByCommentID(c echo.Context) error
	CreateReply(c echo.Context) error
	UpdateReply(c echo.Context) error
	DeleteReply(c echo.Context) error
}

type replyController struct {
	ru usecase.IReplyUseCase
}

func NewReplyController(ru usecase.IReplyUseCase) IReplyController {
	return &replyController{ru}
}

func (rc *replyController) GetAllReplies(c echo.Context) error {
	replies, err := rc.ru.GetAllReplies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, replies)
}

func (rc *replyController) GetAllRepliesById(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	reply, err := rc.ru.GetReplyByID(uint(replyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, reply)
}

func (rc *replyController) GetRepliesByCommentID(c echo.Context) error {
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	replies, err := rc.ru.GetRepliesByCommentID(uint(commentId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, replies)
}

func (rc *replyController) CreateReply(c echo.Context) error {
	reply := model.Reply{}
	if err := c.Bind(&reply); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	replyResponse, err := rc.ru.CreateReply(reply)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, replyResponse)
}

func (rc *replyController) UpdateReply(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	reply := model.Reply{}
	if err := c.Bind(&reply); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	replyResponse, err := rc.ru.UpdateReply(reply, uint(replyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, replyResponse)
}

func (rc *replyController) DeleteReply(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = rc.ru.DeleteReply(uint(replyId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Reply Deleted")
}
