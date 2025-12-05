package http

import (
	"net/http"
	"net/url"
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
		redirectURL, err := url.JoinPath(
			h.cfgSrv.KCBasePath,
			"realms",
			h.cfgSrv.KCRealm,
			strings.TrimPrefix(c.Request.URL.String(), "/auth"),
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, NewError(err.Error()))
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	}
}
