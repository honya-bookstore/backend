package http

import (
	"net/http"
	"strings"

	"backend/config"

	"github.com/gin-gonic/gin"
)

type AuthHandlerImpl struct {
	cfgSrv *config.Server
}

func ProvideAuthHandler(cfg *config.Server) *AuthHandlerImpl {
	return &AuthHandlerImpl{
		cfgSrv: cfg,
	}
}

func (h *AuthHandlerImpl) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := h.cfgSrv.PublicKeycloakURL
		if path == "" {
			path = h.cfgSrv.KCBasePath
		}
		redirectURL := path + strings.TrimPrefix(c.Request.URL.String(), "/auth")
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}
