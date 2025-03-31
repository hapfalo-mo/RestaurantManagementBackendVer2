package interfaces

import (
	dto "RestuarantBackend/models/dto"
)

type FoodInterface interface {
	GetAllFoodPagingList(request *dto.PagingRequest) (result []dto.FoodMenuResponse, err error)
}
