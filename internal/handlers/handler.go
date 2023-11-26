package handlers

import (
	"helloladies/apps/backend/internal/lib/jwt"
	"helloladies/apps/backend/internal/middleware"
	"helloladies/apps/backend/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Services
	log      *logrus.Logger
	cfg      jwt.Config
}

func New(services *service.Services, cfg jwt.Config, log *logrus.Logger) *Handler {
	return &Handler{
		services: services,
		log:      log,
		cfg:      cfg,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(h.CORSMiddleware())
	middlewares := middleware.NewMiddlewares(h.cfg)

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.SignUp)
			auth.POST("/sign-in", h.SignIn)
		}

		logged := api.Group("/logged", middlewares.VerifyToken)
		{
			logged.GET("/info", h.Ping)
		}
	}

	return router
}

func (h *Handler) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		c.Next()
	}
}
