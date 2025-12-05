package http

import (
	"errors"
	"net/http"
	"strings"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hashicorp/go-multierror"
)

type AuthMiddlewareImpl struct {
	keycloakClient      *gocloak.GoCloak
	srvCfg              *config.Server
	missingJWTErr       string
	invalidJWTErr       string
	decodeJWTErr        string
	failedIntrospectErr string
	inactiveTokenErr    string
}

var _ AuthMiddleware = &AuthMiddlewareImpl{}

func ProvideAuthMiddleware(
	keycloakClient *gocloak.GoCloak,
	srvCfg *config.Server,
) *AuthMiddlewareImpl {
	return &AuthMiddlewareImpl{
		keycloakClient:      keycloakClient,
		srvCfg:              srvCfg,
		missingJWTErr:       "Missing Authorization header",
		invalidJWTErr:       "Invalid Authorization header format",
		decodeJWTErr:        "Cannot decode access token",
		failedIntrospectErr: "Failed to introspect token",
		inactiveTokenErr:    "Inactive or invalid token",
	}
}

func (m *AuthMiddlewareImpl) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewError(m.missingJWTErr))
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				NewError(m.invalidJWTErr),
			)
			return
		}
		token := parts[1]
		tokens, _, err := m.keycloakClient.DecodeAccessToken(ctx, token, m.srvCfg.KCRealm)
		if err != nil {
			err = multierror.Append(err, errors.New(m.decodeJWTErr))
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				NewError(err.Error()),
			)
			return
		}
		rptResult, err := m.keycloakClient.RetrospectToken(
			ctx,
			token,
			m.srvCfg.KCClientId,
			m.srvCfg.KCClientSecret,
			m.srvCfg.KCRealm,
		)
		if err != nil {
			err := multierror.Append(err, errors.New(m.failedIntrospectErr))
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				NewError(err.Error()),
			)
			return
		}
		if rptResult == nil || !*rptResult.Active {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				NewError(m.inactiveTokenErr),
			)
			return
		}

		claims, _ := tokens.Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)
		ctx.Set("userID", userID)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
