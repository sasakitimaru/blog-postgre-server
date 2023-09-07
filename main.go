package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	dbConn := db.NewDB()
	articleValidator := validator.NewArticleValidator()
	commentValidator := validator.NewCommentValidator()
	articleRepository := repository.NewArticleRepository(dbConn)
	commentRepository := repository.NewCommentRepository(dbConn)
	replyRepository := repository.NewReplyRepository(dbConn)
	articleUsecase := usecase.NewArticleUseCase(articleRepository, articleValidator)
	commentUseCase := usecase.NewCommentUseCase(commentRepository, commentValidator)
	replyUseCase := usecase.NewReplyUseCase(replyRepository, commentValidator)
	articleController := controller.NewArticleController(articleUsecase)
	commentController := controller.NewCommentController(commentUseCase)
	replyController := controller.NewReplyController(replyUseCase)
	e := router.NewRouter(articleController, commentController, replyController)
	e.Logger.Fatal(e.Start(":8080"))
	db.CloseDB(dbConn)
}
