package handlers

import (
	"RestuarantBackend/interfaces"
	dto "RestuarantBackend/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FoodController struct {
	service interfaces.FoodInterface
}

// Constructor for Dependencies Injection
func NewFoodController(service interfaces.FoodInterface) *FoodController {
	if service == nil {
		panic("NewFoodController is null")
	}
	return &FoodController{service: service}
}

// Get all Food Paging List
func (f *FoodController) GetAllFoodPagingList(c *gin.Context) {
	var request dto.PagingRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Request does not work.Pls try it again!"})
		return
	}
	if f.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Server doesn't work! Pls try it again!"})
		return
	}
	result, err := f.service.GetAllFoodPagingList(&request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Something does not work!"})
		return
	}
	c.JSON(http.StatusOK, result)
}
