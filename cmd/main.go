package main

import (
	"context"
	"io"
	"log"

	"backend/internal/di"

	_ "backend/docs"

	"github.com/gin-gonic/gin"
)

//	@BasePath	/api

//	@securitydefinitions.oauth2.password	OAuth2Password
//	@tokenUrl								/auth/protocol/openid-connect/token

//	@securitydefinitions.oauth2.accessCode	OAuth2AccessCode
//	@tokenUrl								/auth/protocol/openid-connect/token
//	@authorizationUrl						/auth/protocol/openid-connect/auth

func main() {
	ctx := context.Background()
	s := di.InitializeServer(ctx)
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	err := s.Run()
	if err != nil {
		log.Fatal("Server run error", err)
	}
}
