package http

import (
	"backend/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine      *gin.Engine
	srvCfg      *config.Server
	redisClient *redis.Client
	authHandler AuthHandler
}

func NewServer(
	e *gin.Engine,
	r Router,
	srvCfg *config.Server,
	redisClient *redis.Client,
	authHandler AuthHandler,
) *Server {
	e.Use(cors.New(cors.Config{
		AllowOrigins: srvCfg.AllowOrigins,
	}))
	r.RegisterRoutes(e)
	auth := e.Group("/auth")
	{
		auth.GET("/*path", authHandler.Handler())
		auth.POST("/*path", authHandler.Handler())
	}
	e.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(
			swaggerfiles.Handler,
			ginSwagger.PersistAuthorization(true),
			ginSwagger.Oauth2DefaultClientID("swagger"),
		),
	)
	return &Server{
		engine:      e,
		srvCfg:      srvCfg,
		authHandler: authHandler,
	}
}

func (s *Server) Run() error {
	return s.engine.Run()
}
