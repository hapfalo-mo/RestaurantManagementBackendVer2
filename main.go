package main

import (
	"RestuarantBackend/config"
	"RestuarantBackend/db"
	"RestuarantBackend/models/events"
	"RestuarantBackend/routers"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"github.com/hapfalo-mo/RestaurantUserService/restaurantuserservicerpb"
	"log"
)

//func SumTwoNumber(c restaurantuserservicerpb.RestaurantUserServiceClient) {
//	resp, err := c.Sum(context.Background(), &restaurantuserservicerpb.SumRequest{
//		Num1: 5,
//		Num2: 3,
//	})
//	if err != nil {
//		log.Fatalf("Failed to call Sum Function: %v", err)
//	}
//	log.Printf("Sum result: %f ", resp.Result)
//}

func main() {
	// Integrate Kafka Service
	brokers := []string{"localhost:9092"}
	topic := "order-events"
	kafkaClient, err := config.GenerateNewKafkaClient(brokers, topic)
	if err != nil {
		log.Fatalf("Failed to create kafka client: %v", err)
	}
	defer kafkaClient.Close()

	// Test publish event
	PublishOrderService(kafkaClient)
	// Connect Database
	db.Connect()
	defer db.DB.Close()
	// Run client
	//cl := client.RunClient()
	//SumTwoNumber(cl)
	// Initialize Router
	router := gin.Default()

	// CORS Middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:1703"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	// //Xử lý OPTIONS request để tránh bị block bởi preflight request
	// router.OPTIONS("/*path", func(c *gin.Context) {
	// 	c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
	// 	c.Header("Access-Control-Allow-Origin", "http://localhost:1703")
	// 	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
	// 	c.Header("Access-Control-Allow-Credentials", "true")
	// 	c.Status(204) // No Content
	// })

	// Register User API Routes
	routers.SetRoutesAPI(router)

	// Run Server
	router.Run(":1611")
}

func PublishOrderService(kafkaClient *config.KafkaClient) {
	orderMessage := events.OrderMessage{
		OrderId:  "1",
		Message:  "Đặt đơn hàng thành công",
		Amount:   3,
		Money:    500000,
		Username: "Dao Khuat",
	}

	ctx := context.Background()
	err := kafkaClient.PublistOrderService(ctx, orderMessage)
	if err != nil {
		log.Println("Failed to publish order event:", err)
	} else {
		log.Println("Order event published successfully:", orderMessage)
	}
}
