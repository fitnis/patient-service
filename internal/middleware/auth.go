package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fitnis/patient-service/configs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// JWTConfig stores JWT middleware configuration
type JWTConfig struct {
	KeycloakBaseURL  string
	KeycloakRealm    string
	KeycloakClientID string
}

// NewJWTConfig creates a new JWTConfig
func NewJWTConfig(config *configs.Config) *JWTConfig {
	return &JWTConfig{
		KeycloakBaseURL:  config.KeycloakBaseURL,
		KeycloakRealm:    config.KeycloakRealm,
		KeycloakClientID: config.KeycloakClientID,
	}
}

// JWTMiddleware returns a middleware that validates JWT tokens
func JWTMiddleware(config *JWTConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the token from the Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing authorization header")
			}

			// The token should be in the format "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid authorization header format")
			}

			tokenString := parts[1]

			// Parse the token without verification (we'll verify with Keycloak introspection in a real implementation)
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// In a real implementation, you would verify the signature by getting the key from Keycloak
				// For simplicity, we're just checking if the token exists and has claims
				return nil, nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("invalid token: %v", err))
			}

			// Store the token claims in the context for later use
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("user", claims)
			}

			return next(c)
		}
	}
}

// HasRole checks if the current user has the required role
func HasRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userClaims, ok := c.Get("user").(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "user claims not found")
			}

			// Extract roles from the token
			// In a Keycloak JWT, roles are typically in realm_access.roles
			realmAccess, ok := userClaims["realm_access"].(map[string]interface{})
			if !ok {
				return echo.NewHTTPError(http.StatusForbidden, "no realm access information")
			}

			roles, ok := realmAccess["roles"].([]interface{})
			if !ok {
				return echo.NewHTTPError(http.StatusForbidden, "no roles information")
			}

			// Check if the user has the required role
			for _, r := range roles {
				if r == role {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "insufficient permissions")
		}
	}
}
