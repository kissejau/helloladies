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

const (
	errIncorrectBody = "incorrect body"
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
			logged.GET("/info", h.Ping)

			users := logged.Group("/users")
			{
				users.GET("/get", h.GetUser)
				users.GET("/all", h.ListUsers)
				users.PUT("/update", h.UpdateUser)
				users.DELETE("/delete", h.DeleteUser)
			}

			cities := logged.Group("/cities")
			{
				cities.GET("/all", h.ListCities)

				admin := cities.Group("/", middlewares.VerifyAdminPermissions)
				{
					admin.PUT("/update", h.UpdateCity)
					admin.POST("/create", h.CreateCity)
					admin.DELETE("/delete", h.DeleteCity)
				}
			}

			univs := logged.Group("/univs")
			{
				univs.GET("/all", h.ListUnivs)
				univs.GET("/list", h.GetUnivsByCity) // by city code

				admin := univs.Group("/", middlewares.VerifyAdminPermissions)
				{
					admin.PUT("/update", h.UpdateUniv)
					admin.POST("/create", h.CreateUniv)
					admin.DELETE("/delete", h.DeleteUniv)
				}
			}

			teachers := logged.Group("/teachers")
			{
				teachers.GET("/all", h.ListTeachers)
				teachers.GET("/list", h.GetTeacherByUniv)

				admin := teachers.Group("/", middlewares.VerifyAdminPermissions)
				{
					admin.PUT("/update", h.UpdateTeacher)
					admin.POST("/create", h.CreateTeacher)
					admin.DELETE("/delete", h.DeleteTeacher)
				}
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
