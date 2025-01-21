package main

import (
	proto "grpc/proto/gen/go"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Подключение к gRPC серверу
	connection, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer connection.Close()

	// Инициализация клиента
	c := proto.NewSumServiceClient(connection)

	// Контекст с таймаутом в одну секунду
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Вызов удаленного метода суммирования двуз чисел
	result, err := c.Sum(ctx, &proto.SumRequest{N1: 2, N2: 2})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// Вывести результат в консоль
	log.Println(result)
}
