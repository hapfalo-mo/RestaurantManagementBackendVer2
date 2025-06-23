package interfaces

import (
	custom "RestuarantBackend/custom"
	dto "RestuarantBackend/models/dto"
)

type OrderInterface interface {
	CreateNewOrder(request *dto.OrderCreateRequest) (result custom.Data[dto.OrderResponse], err custom.Error)
	CreateOrderItems(request *dto.OrderItemRequest) (result string, err error)
	GetOrderById(id int) (result custom.Data[[]dto.OrderDetailResponse], err custom.Error)
	GetAllOrderByUserId(userId int, request *dto.PagingRequest) (result custom.Data[[]dto.OrderResponse], err custom.Error)
	GenerateAndConfirmOTP(userId int, userEmail string) (result bool, err custom.Error)
}
