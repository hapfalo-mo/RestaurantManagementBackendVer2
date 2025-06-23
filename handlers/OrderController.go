package handlers

import (
	"RestuarantBackend/custom"
	errorList "RestuarantBackend/error"
	"RestuarantBackend/interfaces"
	"fmt"
	"net/http"
	"strconv"

	dto "RestuarantBackend/models/dto"

	"github.com/gin-gonic/gin"
)

type OrderControlloer struct {
	service interfaces.OrderInterface
}

// Constructor for denpendency injection
func NewOrderController(service interfaces.OrderInterface) *OrderControlloer {
	if service == nil {
		panic("NewOrderontroller service is nil")
	}
	return &OrderControlloer{service: service}
}

// Create new Order
func (o *OrderControlloer) CreateNewOrder(c *gin.Context) {
	var request dto.OrderCreateRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom.Error{
			Message:    errorList.ErrBadRequest.Error(),
			ErrorField: err.Error(),
			Field:      "OrderController - CreateNewOrder",
		})
		return
	}
	if o.service == nil {
		c.JSON(http.StatusInternalServerError, custom.Error{
			Message:    errorList.ErrInternalServer.Error(),
			ErrorField: err.Error(),
			Field:      "OrderController - CreateNewOrder",
		})
		return
	}
	result, errResponse := o.service.CreateNewOrder(&request)
	if errResponse.Message != "" {
		c.JSON(http.StatusNotFound, errResponse)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Create new Order
func (o *OrderControlloer) CreateOrderItems(c *gin.Context) {
	var request dto.OrderItemRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom.Error{
			Message:    errorList.ErrBadRequest.Error(),
			ErrorField: err.Error(),
			Field:      "OrderController - CreateOrderItems",
		})
	}
	if o.service == nil {
		c.JSON(http.StatusInternalServerError, custom.Error{
			Message:    errorList.ErrInternalServer.Error(),
			ErrorField: err.Error(),
			Field:      "OrderController - CreateNewOrder",
		})
		return
	}
	result, errorResponse := o.service.CreateOrderItems(&request)
	if errorResponse != nil {
		c.JSON(http.StatusNotFound, errorResponse)
		return
	}
	c.JSON(http.StatusOK, result)
}

// Get All Order By UserId
func (o *OrderControlloer) GetAllOrderByUserId(c *gin.Context) {
	var request dto.PagingRequest
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, custom.Error{
			Message:    errorList.ErrBadRequest.Error(),
			ErrorField: err.Error(),
			Field:      "OrderController - GetAllOrderByUserId",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Cannot convert Id from string to int"})
		return
	}
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, custom.Error{
			Message:    errorList.ErrBadRequest.Error(),
			ErrorField: err.Error(),
			Field:      "OrderController - GetAllOrderByUserId",
		})
		return
	}
	result, errorResponse := o.service.GetAllOrderByUserId(id, &request)
	if errorResponse.Message != "" {
		c.JSON(http.StatusNotFound, errorResponse)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": result})
}

// Get OrderDetail By OrderId
func (o *OrderControlloer) GetOrderById(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Println(idStr)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid order ID"})
		return
	}
	result, errorResponse := o.service.GetOrderById(id)
	if errorResponse.Message != "" {
		c.JSON(http.StatusNotFound, errorResponse)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Send OTP

func (o *OrderControlloer) GenerateOTP(c *gin.Context) {
	idStr := c.Param("id")
	userEmail := c.Param("userEmail")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid order ID"})
		return
	}
	_, errResponse := o.service.GenerateAndConfirmOTP(id, userEmail)
	if errResponse.Message != "" {
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Gửi mã OTP thành công"})
}
