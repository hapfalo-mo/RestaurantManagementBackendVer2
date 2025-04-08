package routers

import (
	db "RestuarantBackend/db"
	handlers "RestuarantBackend/handlers"
	middleware "RestuarantBackend/middleware"
	service "RestuarantBackend/service"

	"github.com/gin-gonic/gin"
)

func SetRoutesAPI(r *gin.Engine) {
	db.Connect()
	userService := &service.UserService{}
	bookingService := &service.BookingService{}
	foodService := &service.FoodService{}
	orderService := &service.OrderService{}
	bookingController := handlers.NewBookingController(bookingService)
	userController := handlers.NewUserController(userService)
	foodController := handlers.NewFoodController(foodService)
	orderController := handlers.NewOrderController(orderService)
	v1 := r.Group("api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/signup", userController.Register)
			users.POST("/login", userController.LoginToken)
			users.POST("/login-googgle", userController.LoginGoogle)
			users.PUT("/updateUser", middleware.AuthenticateMiddleware, userController.Update)
			users.POST("/getAllUser", middleware.AuthenAdminMiddelWare, userController.GetAllUSerPagingList)
			users.POST("/getAllUserVer2", middleware.AuthenticateMiddleware, userController.GetAllUser)
			users.GET("/export-csvFile", middleware.AuthenAdminMiddelWare, userController.ExportUserCSVFile)
			users.PUT("/block-unblock-user/:id", middleware.AuthenAdminMiddelWare, userController.BlockOrUnblockUser)
		}
		bookings := v1.Group("/bookings")
		{
			bookings.POST("/bookTable", middleware.AuthenticateMiddleware, bookingController.BookingTable)
			bookings.POST("/getBooking/:id", middleware.AuthenticateMiddleware, bookingController.PagingBookingList)
			bookings.POST("/get-all-bookings", middleware.AuthenAdminMiddelWare, bookingController.PagingAllBookingList)

		}
		foods := v1.Group("/foods")
		{
			foods.POST("/get-all-foods", foodController.GetAllFoodPagingList)
		}
		orders := v1.Group("/orders")
		{
			orders.POST("/create-order", middleware.AuthenticateMiddleware, orderController.CreateNewOrder)
			orders.POST("/create-order-items", middleware.AuthenticateMiddleware, orderController.CreateOrderItems)
			orders.POST("/get-all-order/:id", middleware.AuthenticateMiddleware, orderController.GetAllOrderByUserId)
			orders.GET("get-order-detail/:id", middleware.AuthenticateMiddleware, orderController.GetOrderById)
		}
	}
}
