package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/ngdangkietse/ndk-go-proto/generated/account"
	"github.com/ngdangkietse/ndk-go-proto/generated/common"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := account.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.UpsertUser(ctx, &account.PUser{
		Id:       uuid.New().String(),
		FullName: "NgDangKiet",
		Email:    "kiet.nguyen-dang@dev.com",
	})

	response2, err := client.FindUserByEmail(ctx, &common.PEmailRequest{
		Email: "ngdangkiet@yopmail.com",
	})

	if err != nil {
		log.Fatalf("Count not upsert user: %v", err)
	}

	log.Println(response.Message)
	log.Println(response2.Data)
}
