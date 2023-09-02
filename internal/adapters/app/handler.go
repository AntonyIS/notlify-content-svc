package app

import (
	"fmt"

	appConfig "github.com/AntonyIS/notlify-content-svc/config"
	"github.com/AntonyIS/notlify-content-svc/internal/adapters/logger"
	"github.com/AntonyIS/notlify-content-svc/internal/core/ports"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitGinRoutes(svc ports.ArticleService, logger logger.LoggerType, conf appConfig.Config) {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8001"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	handler := NewGinHandler(svc, conf.SECRET_KEY)

	articleRoutes := router.Group("/v1/articles")
	{
		articleRoutes.POST("/", handler.CreateArticle)
		articleRoutes.GET("/:article_id", handler.GetArticleByID)
		articleRoutes.GET("/", handler.GetArticles)
		articleRoutes.GET("/author/:author_id", handler.GetArticlesByAuthor)
		articleRoutes.GET("/tag/:tag_name", handler.GetArticlesByTag)
		articleRoutes.PUT("/", handler.UpdateArticle)
		articleRoutes.DELETE("/:id", handler.DeleteArticle)
		articleRoutes.DELETE("/", handler.DeleteArticleAll)
	}

	logger.PostLogMessage(fmt.Sprintf("Server running on port :%s", conf.Port))
	router.Run(fmt.Sprintf(":%s", conf.Port))
}