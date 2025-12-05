package http

import (
	"net/http"

	"backend/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type RoleMiddlewareImpl struct {
	noJWTFoundErr              string
	invalidJWTErr              string
	noRoleFoundErr             string
	insufficientPermissionsErr string
	srvCfg                     *config.Server
}

var _ RoleMiddleware = &RoleMiddlewareImpl{}

func ProvideRoleMiddleware(srvCfg *config.Server) *RoleMiddlewareImpl {
	return &RoleMiddlewareImpl{
		srvCfg:                     srvCfg,
		noJWTFoundErr:              "no JWT token found",
		invalidJWTErr:              "invalid JWT token",
		noRoleFoundErr:             "no role found in JWT claims",
		insufficientPermissionsErr: "insufficient permissions",
	}
}

func (m *RoleMiddlewareImpl) Handler(rolesAllowed []UserRole) gin.HandlerFunc {
	set := make(map[UserRole]struct{})
	for _, requiredRole := range rolesAllowed {
		set[requiredRole] = struct{}{}
	}
	return func(ctx *gin.Context) {
		claimsInterface, exists := ctx.Get("claims")
		if !exists {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				NewError(m.noJWTFoundErr))
			return
		}

		claims, ok := claimsInterface.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, NewError(m.invalidJWTErr))
			return
		}

		userRoles := extractRole(claims)
		if len(userRoles) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(m.noRoleFoundErr))
			return
		}

		allowed := false
		for _, role := range userRoles {
			if _, exists := set[UserRole(role)]; exists {
				ctx.Set("user_role", role)
				allowed = true
				break
			}
		}
		if !allowed {
			ctx.AbortWithStatusJSON(http.StatusForbidden, NewError(m.insufficientPermissionsErr))
			return
		}

		ctx.Next()
	}
}

func extractRole(claims jwt.MapClaims) []string {
	roles := []string{}
	if realmAccess, ok := claims["realm_access"].(map[string]interface{}); ok {
		if realmRoles, ok := realmAccess["roles"].([]interface{}); ok {
			for _, r := range realmRoles {
				if roleStr, ok := r.(string); ok {
					roles = append(roles, roleStr)
				}
			}
		}
	}
	return roles
}
