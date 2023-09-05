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
	articleRepository := repository.NewArticleRepository(dbConn)
	articleUsecase := usecase.NewArticleUseCase(articleRepository, articleValidator)
	articleController := controller.NewArticleController(articleUsecase)
	e := router.NewRouter(articleController)
	e.Logger.Fatal(e.Start(":8080"))
	db.CloseDB(dbConn)
}
