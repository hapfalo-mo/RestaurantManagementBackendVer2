package service

import (
	"RestuarantBackend/db"
	"RestuarantBackend/interfaces"
	dto "RestuarantBackend/models/dto"
)

var _ interfaces.FoodInterface = &FoodService{}

type FoodService struct{}

func (f *FoodService) GetAllFoodPagingList(request *dto.PagingRequest) (result []dto.FoodMenuResponse, err error) {
	offset := (request.Page - 1) * request.PageSize
	prepareStatement := "SELECT id, name, price, description, image_url, created_at, updated_at, deleted_at FROM food LIMIT ? OFFSET ?"
	rows, err := db.DB.Query(prepareStatement, request.PageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item dto.FoodMenuResponse
		err = rows.Scan(&item.Id, &item.FoodName, &item.Price, &item.Description, &item.ImageURL, &item.CreatedAt, &item.UpdatedAt, &item.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil
}
