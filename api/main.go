package main

import (
	"context"
	"grpc/proto/gen/go"
	"log"
	"net"

	"google.golang.org/grpc"
)

// SumServiceServer структура реализует интерфейс сгенерированный компилятором прото файла.
type SumServiceServer struct {
	proto.UnimplementedSumServiceServer
}

// Sum - реализует метод в структуре SumServiceServer.
// Метод возвращает сумму двух чисел из gRPC запроса.
func (s *SumServiceServer) Sum(ctx context.Context, request *proto.SumRequest) (*proto.SumResponse, error) {
	log.Default().Printf("sum %v + %v", request.N1, request.N2)
	return &proto.SumResponse{Sum: request.N1 + request.N2}, nil
}

func main() {
	// Прослушивание локальной сети на входящие запросы.
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Инициализация gRPC сервера.
	grpcServer := grpc.NewServer()
	proto.RegisterSumServiceServer(grpcServer, &SumServiceServer{})

	// Запуск gRPC сервера, который принимает входящие запросы и обрабатывает их.
	log.Printf("server listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
