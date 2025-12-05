package http

import (
	"backend/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	engine      *gin.Engine
	srvCfg      *config.Server
	authHandler AuthHandler
}

func NewServer(
	e *gin.Engine,
	r Router,
	srvCfg *config.Server,
	authHandler AuthHandler,
) *Server {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = srvCfg.AllowOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	e.Use(cors.New(corsConfig))

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
