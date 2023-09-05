package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IArticleUseCase interface {
	GetAllArticles() ([]model.ArticleResponse, error)
	GetArticleByID(articleId uint) (model.ArticleResponse, error)
	CreateArticle(article model.Article) (model.ArticleResponse, error)
	UpdateArticle(article model.Article, articleId uint) (model.ArticleResponse, error)
	DeleteArticle(articleId uint) error
}

type articleUseCase struct {
	ar repository.IArticleRepository
	av validator.IArticleValidator
}

func NewArticleUseCase(ar repository.IArticleRepository, av validator.IArticleValidator) IArticleUseCase {
	return &articleUseCase{ar, av}
}

func (au *articleUseCase) GetAllArticles() ([]model.ArticleResponse, error) {
	articles := []model.Article{}
	articleResponse := []model.ArticleResponse{}
	if err := au.ar.GetAllArticles(&articles); err != nil {
		return nil, err
	}
	for _, article := range articles {
		a := model.ArticleResponse{
			ID:    article.ID,
			Title: article.Title,
			Likes: article.Likes,
		}
		articleResponse = append(articleResponse, a)

	}
	return articleResponse, nil
}

func (au *articleUseCase) GetArticleByID(articleId uint) (model.ArticleResponse, error) {
	article := model.Article{}
	if err := au.ar.GetArticleByID(&article, articleId); err != nil {
		return model.ArticleResponse{}, err
	}
	articleResponse := model.ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Likes: article.Likes,
	}
	return articleResponse, nil
}

func (au *articleUseCase) CreateArticle(article model.Article) (model.ArticleResponse, error) {
	if err := au.av.ArticleValidate(article); err != nil {
		return model.ArticleResponse{}, err
	}
	if err := au.ar.CreateArticle(&article); err != nil {
		return model.ArticleResponse{}, err
	}
	articleResponse := model.ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Likes: article.Likes,
	}
	return articleResponse, nil
}

func (au *articleUseCase) UpdateArticle(article model.Article, articleId uint) (model.ArticleResponse, error) {
	if err := au.av.ArticleValidate(article); err != nil {
		return model.ArticleResponse{}, err
	}
	if err := au.ar.UpdateArticle(&article, articleId); err != nil {
		return model.ArticleResponse{}, err
	}
	articleResponse := model.ArticleResponse{
		ID:    article.ID,
		Title: article.Title,
		Likes: article.Likes,
	}
	return articleResponse, nil
}

func (au *articleUseCase) DeleteArticle(articleId uint) error {
	if err := au.ar.DeleteArticle(articleId); err != nil {
		return err
	}
	return nil
}
