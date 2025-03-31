package handlers

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad Request"})
		return
	}
	if o.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Error"})
		return
	}
	result, err := o.service.CreateNewOrder(&request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Wrong in processing.."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Data": result})
}

// Create new Order
func (o *OrderControlloer) CreateOrderItems(c *gin.Context) {
	var request dto.OrderItemRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad Request"})
		return
	}
	if o.service == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Error"})
		return
	}
	result, err := o.service.CreateOrderItems(&request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Wrong in processing.."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": result})
}

// Get All Order By UserId
func (o *OrderControlloer) GetAllOrderByUserId(c *gin.Context) {
	var request dto.PagingRequest
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid order ID"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Cannot convert Id from string to int"})
		return
	}
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad Request"})
		return
	}
	result, err := o.service.GetAllOrderByUserId(id, &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Have error in get orders"})
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
	result, err := o.service.GetOrderById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order details"})
		return
	}

	c.JSON(http.StatusOK, result)
}
