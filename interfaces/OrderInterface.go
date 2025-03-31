package interfaces

import (
	dto "RestuarantBackend/models/dto"
)

type OrderInterface interface {
	CreateNewOrder(request *dto.OrderCreateRequest) (result *dto.OrderResponse, err error)
	CreateOrderItems(request *dto.OrderItemRequest) (result string, err error)
	GetOrderById(id int) (result []dto.OrderDetailResponse, err error)
	GetAllOrderByUserId(userId int, request *dto.PagingRequest) (result []dto.OrderResponse, err error)
}
