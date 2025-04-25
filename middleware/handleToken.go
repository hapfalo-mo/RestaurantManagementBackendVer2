package middleware

import (
	"RestuarantBackend/custom"
	errorList "RestuarantBackend/error"
	service "RestuarantBackend/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticateMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		var response = custom.Error{
			Message:    errorList.ErrCreatingToken.Error(),
			ErrorField: "Empty AuthHeader",
			Field:      "Authen-Token",
		}
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		var response2 = custom.Error{
			Message:    errorList.ErrInvalidToken.Error(),
			ErrorField: "Invalid Token",
			Field:      "Authen-Token",
		}
		c.JSON(http.StatusUnauthorized, response2)
		c.Abort()
		return
	}
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
		return
	}
	ok, err := service.CallApiCheckUser(tokenString)
	if err != nil || ok == false {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Can not verify user!"})
		c.Abort()
		return
	}
	c.Next()
}

// Function for Admin Authen and Authorize
func AuthenAdminMiddelWare(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Missing Token"})
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "Invalid Token Format"})
		c.Abort()
		return
	}
	// Parse Token to get Claim
	claims, err := service.ParseToken(tokenString)
	if err != nil {
		c.JSON(401, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	if claims.Role != 1 {
		c.JSON(http.StatusForbidden, gin.H{"Error": "You don't have permission to access this site! "})
		c.Abort()
		return
	}
	c.Next()
}
