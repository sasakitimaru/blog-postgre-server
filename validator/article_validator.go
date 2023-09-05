package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IArticleValidator interface {
	ArticleValidate(article model.Article) error
}

type articleValidator struct{}

func NewArticleValidator() IArticleValidator {
	return &articleValidator{}
}
func (av *articleValidator) ArticleValidate(article model.Article) error {
	return validation.ValidateStruct(&article,
		validation.Field(&article.Title, validation.Required.Error("title is required"), validation.Length(5, 40).Error("limited min 10 to max 40 characters")),
	)
}
