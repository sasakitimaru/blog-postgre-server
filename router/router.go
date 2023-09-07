package router

import (
	"go-rest-api/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(ac controller.IArticleInterface, cc controller.ICommentController, rc controller.IReplyController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowCredentials: true,
	}))
	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	CookiePath: "/",
	// 	CookieDomain:   os.Getenv("API_DOMAIN"),
	// 	CookieHTTPOnly: true,
	// 	// When you are using localhost, you should set SameSiteDefaultMode.
	// 	CookieSameSite: http.SameSiteDefaultMode,
	// 	// When you are using production, you should set SameSiteNoneMode.
	// 	// CookieSameSite: http.SameSiteNoneMode,
	// }))
	e.GET("/csrf", ac.CsrfToken)
	t := e.Group("/articles")
	t.GET("", ac.GetAllArticles)
	t.GET("/:articleId", ac.GetAllArticlesById)
	t.POST("", ac.CreateArticle)
	t.PUT("/:articleId", ac.UpdateArticle)
	t.DELETE("/:articleId", ac.DeleteArticle)

	c := e.Group("/comments")
	c.GET("", cc.GetAllComments)
	c.GET("/:commentId", cc.GetAllCommentsById)
	c.GET("/article/:articleId", cc.GetCommentsByArticleID)
	c.POST("", cc.CreateComment)
	c.PUT("/:commentId", cc.UpdateComment)
	c.DELETE("/:commentId", cc.DeleteComment)

	r := e.Group("/replies")
	r.GET("", rc.GetAllReplies)
	r.GET("/:replyId", rc.GetAllRepliesById)
	r.GET("/comment/:commentId", rc.GetRepliesByCommentID)
	r.POST("", rc.CreateReply)
	r.PUT("/:replyId", rc.UpdateReply)
	r.DELETE("/:replyId", rc.DeleteReply)
	return e
}
