package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/hapfalo-mo/RestaurantUserService/restaurantuserservicerpb"
)

type AuthController struct{}

func (au *AuthController) AuthUserHandler(c *gin.Context, us restaurantuserservicerpb.RestaurantUserServiceClient) (bool, error) {
	token, err := c.Cookie("token")
	if err != nil {
		return false, err
	}
	req := &restaurantuserservicerpb.IsVerifyUserRequest{
		Token: token,
	}
	_, err = us.IsAcceptUserAccess(context.Background(), req)
	if err != nil {
		return false, err
	}
	return true, nil
}
