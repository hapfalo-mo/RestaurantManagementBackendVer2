package client

import (
	"github.com/hapfalo-mo/RestaurantUserService/restaurantuserservicerpb"
	"google.golang.org/grpc"
	"log"
)

func RunClient() restaurantuserservicerpb.RestaurantUserServiceClient {
	cc, err := grpc.Dial("localhost:50016", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dialing : %v", err)
	}
	client := restaurantuserservicerpb.NewRestaurantUserServiceClient(cc)
	return client
}
