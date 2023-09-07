package validator

import (
	"errors"
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICommentValidator interface {
	CommentValidate(comment model.Comment, commentsForValidate *[]model.Comment) error
	CommentUpdateValidate(comment model.Comment, TargetCommentAuthor string) error
}

type commentValidator struct{}

func NewCommentValidator() ICommentValidator {
	return &commentValidator{}
}

func (cv *commentValidator) CommentValidate(comment model.Comment, commentsForValidate *[]model.Comment) error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.Comment, validation.Required.Error("comment is required"), validation.Length(5, 100).Error("limited min 5 to max 100 characters")),
		validation.Field(&comment.Author, validation.By(func(value interface{}) error {
			author, ok := value.(string)
			if !ok {
				return errors.New("provided value is not a string")
			}
			if author != "Anonymous" {
				for _, comment := range *commentsForValidate {
					if comment.Author == author {
						return errors.New("author already commented")
					}
				}
			}
			return nil
		})),
	)
}

func (cv *commentValidator) CommentUpdateValidate(comment model.Comment, TargetCommentAuthor string) error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.Author, validation.By(func(value interface{}) error {
			_, ok := value.(string)
			if !ok {
				return errors.New("provided value is not a string")
			}
			// if author != TargetCommentAuthor {
			// 	return errors.New("author does not match target author")
			// }
			return nil
		})),
	)
}
