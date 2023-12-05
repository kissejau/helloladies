package handlers

import (
	"helloladies/internal/lib/jwt"
	"helloladies/internal/middleware"
	"helloladies/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	middlewares := middleware.NewMiddlewares(h.cfg, h.services.UsersService)

	api := router.Group("/api")
	{
		api.GET("/info", h.Ping)

		router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.SignUp)
			auth.POST("/sign-in", h.SignIn)
		}

		logged := api.Group("/logged", middlewares.VerifyToken)
		{
			users := logged.Group("/users")
			{
				users.GET("/get", h.GetUser)
				users.GET("/all", h.ListUsers)
				users.PUT("/update", h.UpdateUser)
				users.DELETE("/delete", h.DeleteUser)
			}
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
