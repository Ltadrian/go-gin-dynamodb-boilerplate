package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ltadrian/go-gin-dynamodb-boilerplate/config"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := config.GetConfig()
		secret := config.GetString("JWT_SECRET")
		jwtString := c.Request.Header.Get("Authorization")
		if jwtString == "" {
			c.AbortWithStatus(401)
			return
		}

		splitArray := strings.Split(jwtString, "Bearer ")
		tokenString := splitArray[1]

		// validate secret
		type CustomClaims struct {
			Exp int64 `json:"exp"`
			Iss int64 `json:"iss"`
			jwt.StandardClaims
		}
		jwt, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			fmt.Printf("err: %v", err)
			c.AbortWithStatus(500)
			return
		}
		if !jwt.Valid {
			c.AbortWithStatus(401)
			return
		}

		// validate expired token
		if claims, ok := jwt.Claims.(*CustomClaims); ok {
			if isExpired(time.UnixMilli(claims.Exp)) {
				c.AbortWithStatus(401)
				return
			}
			c.Next()
			return
		}
		fmt.Printf("err: %v", err)
		c.AbortWithStatus(500)
	}
}

func isExpired(exp time.Time) bool {
	return time.Now().After(exp)
}
