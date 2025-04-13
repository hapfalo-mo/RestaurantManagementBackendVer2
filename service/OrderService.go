package service

import (
	"RestuarantBackend/custom"
	"RestuarantBackend/db"
	errorList "RestuarantBackend/error"
	"RestuarantBackend/interfaces"
	dto "RestuarantBackend/models/dto"
)

var _ interfaces.OrderInterface = &OrderService{}

type OrderService struct{}

func (o OrderService) CreateNewOrder(request *dto.OrderCreateRequest) (result custom.Data[dto.OrderResponse], errorResponse custom.Error) {

	// Prepare Query Statement
	querry := "INSERT orders (user_id, total_price) VALUES (?,?)"
	res, err := db.DB.Exec(querry, request.UserId, request.TotalPrice)
	if err != nil {
		return custom.Data[dto.OrderResponse]{}, custom.Error{
			Message:    errorList.ErrOrderInsertValue.Error(),
			ErrorField: err.Error(),
			Field:      "OrderService - Querry Database",
		}
	}
	lasInsertId, err := res.LastInsertId()
	if err != nil {
		return custom.Data[dto.OrderResponse]{}, custom.Error{
			Message:    errorList.ErrOrderInsertValue.Error(),
			ErrorField: err.Error(),
			Field:      "OrderService - Querry Database",
		}
	}
	result = custom.Data[dto.OrderResponse]{
		Data: dto.OrderResponse{
			Id: int(lasInsertId),
		}}
	return result, custom.Error{}
}

func (o OrderService) CreateOrderItems(request *dto.OrderItemRequest) (result string, err error) {
	prepareStatement := "INSERT INTO order_items (order_id,food_id,quantity,price) VALUES (?,?,?,?)"
	_, err = db.DB.Exec(prepareStatement, request.OrderId, request.FoodId, request.Quantity, request.Price)
	if err != nil {
		return "", err
	}
	return "Successfully add OrderItem", nil
}

func (o OrderService) GetAllOrderByUserId(userId int, request *dto.PagingRequest) (result custom.Data[[]dto.OrderResponse], errorResonse custom.Error) {
	// Define offset
	offset := (request.Page - 1) * request.PageSize

	// prepareStatement
	querry := "SELECT id, order_status, total_price, ordered_at, updated_at, deleted_at, note, feedback FROM orders WHERE user_id = ? LIMIT ? OFFSET ?"
	rows, err := db.DB.Query(querry, userId, request.PageSize, offset)
	if err != nil {
		return custom.Data[[]dto.OrderResponse]{}, custom.Error{
			Message:    errorList.ErrOrderGetllOrder.Error(),
			ErrorField: err.Error(),
			Field:      "OrderService - Get All Order By UserId",
		}
	}
	defer rows.Close()

	for rows.Next() {
		var response dto.OrderResponse
		err := rows.Scan(&response.Id, &response.OrderStatus, &response.TotalPrice, &response.OrderedAt, &response.UpdatedAt, &response.DeletedAt, &response.Note, &response.Feedback)
		if err != nil {
			return custom.Data[[]dto.OrderResponse]{}, custom.Error{
				Message:    errorList.ErrOrderGetAllOrderScan.Error(),
				ErrorField: err.Error(),
				Field:      "OrderService - Get All Order By UserId - Scan Value",
			}
		}
		result.Data = append(result.Data, response)
	}
	return result, custom.Error{}
}

func (o OrderService) GetOrderById(id int) (result custom.Data[[]dto.OrderDetailResponse], errorResponse custom.Error) {
	// Prepare statement
	querry := `SELECT ot.id, f.name, f.price, ot.quantity, ot.price 
          FROM order_items ot 
          JOIN food f ON ot.food_id = f.id
          WHERE ot.order_id = ?`
	rows, err := db.DB.Query(querry, id)
	if err != nil {
		return custom.Data[[]dto.OrderDetailResponse]{}, custom.Error{
			Message:    errorList.ErrOrderGetAllOrderScan.Error(),
			ErrorField: err.Error(),
			Field:      "OrderService - Get Order By UserId - Scan Value",
		}
	}
	defer rows.Close()
	for rows.Next() {
		var item dto.OrderDetailResponse
		err := rows.Scan(&item.ID, &item.Name, &item.Price, &item.Quantity, &item.Price)
		if err != nil {
			return custom.Data[[]dto.OrderDetailResponse]{}, custom.Error{
				Message:    errorList.ErrOrderGetAllOrderScan.Error(),
				ErrorField: err.Error(),
				Field:      "OrderService - Get Order By UserId - Scan Value",
			}
		}
		result.Data = append(result.Data, item)
	}
	return result, custom.Error{}
}
