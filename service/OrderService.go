package service

import (
	config "RestuarantBackend/config"
	"RestuarantBackend/custom"
	"RestuarantBackend/db"
	errorList "RestuarantBackend/error"
	"RestuarantBackend/interfaces"
	dto "RestuarantBackend/models/dto"
	"context"
	"crypto/rand"
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"math/big"
	"strconv"
	"time"
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

func (o OrderService) GenerateAndConfirmOTP(userEmail string) (result bool, err custom.Error) {

	// Step 1 - Generate OTP
	otp := CreateOTPCode()

	// Step 2 - Save into Redis with Main Key is userEmail
	redisRepo := config.NewRedisClient()
	res := redisRepo.Set(context.Background(), userEmail, otp, 2*time.Minute).Err()
	fmt.Println("Set OTP to Redis", userEmail, otp)
	if res != nil {
		return false, custom.Error{
			Message:    "Error in GenerateAndConfirmOTP",
			ErrorField: res.Error(),
			Field:      "OrderService",
		}
		log.Printf("Error in GenerateAndConfirmOTP", time.Now(), res.Error())
	}

	// Step 3 - Send to Email
	res = SendOTPToEmail(otp, userEmail)
	if res != nil {
		return false, custom.Error{
			Message:    "Error in SendOTPToEmail",
			ErrorField: res.Error(),
			Field:      "OrderService",
		}
		log.Printf("Error in SendOTPToEmail", time.Now(), res.Error())
	}
	return true, custom.Error{}
}

func (o OrderService) IsValidOTP(request dto.OTPRequest) (bool, custom.Error) {
	redisClient := config.NewRedisClient()

	// Check OTP Valid
	isValid := IsEqualOTP(request.UserEmail, request.OTP)
	if !isValid {
		return false, custom.Error{
			Message:    "Mã OTP không hợp lệ",
			ErrorField: "False",
			Field:      "OrderService - IsValidOTP",
		}
		log.Printf("Error in IsValidOTP", time.Now(), "OrderService - IsValidOTP")
	}
	// Remote Key-Value into Redis
	deleted, err := redisClient.Del(context.Background(), request.UserEmail).Result()
	if err != nil {
		return false, custom.Error{}
		log.Printf("Error in IsValidOTP - OrderService", time.Now(), "Cannot Remove Value into Redis")
	}
	if deleted == 0 {
		log.Printf("Warning: OTP key not found in Redis for %s", request.UserEmail)
	}
	return true, custom.Error{}
}

// Domestic Function
func CreateOTPCode() string {
	min := 100000
	max := 999999
	num, err := rand.Int(rand.Reader, big.NewInt(int64(max-min+1)))
	if err != nil {
		return ""
		log.Printf("Error in CreateOTPCode - OrderService.go", time.Now(), err.Error())
	}
	otp := int(num.Int64()) + min
	return strconv.Itoa(otp)
}

func SendOTPToEmail(otp string, userEmail string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "chuongnguyen16112002@gmail.com")
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "OTP SteakHouse")
	m.SetBody("text/plain", "Mã của bạn là: "+otp)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		"chuongnguyen16112002@gmail.com",
		"zapsbztwesnstlww",
	)
	return d.DialAndSend(m)
}

func IsEqualOTP(userEmail string, otp string) bool {

	redis := config.NewRedisClient()
	// Check OTP FROM KEY AND VALUE
	val, err := redis.Get(context.Background(), userEmail).Result()
	if err != nil {
		return false
	}
	if val != otp {
		return false
	} else {
		return true
	}
}
