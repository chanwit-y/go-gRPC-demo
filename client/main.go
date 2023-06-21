package main

import (
	"context"
	"grpc-client/service"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := service.NewLongTimeRequestServiceClient(conn)

	// req := &service.Request{Data: "Test data"}
	stream, err := client.LongTimeRequestStream(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongTimeRequestStream: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break // The server closed the stream.
		}
		if err != nil {
			log.Fatalf("Error while receiving stream: %v", err)
		}
		log.Printf("Response from server: %s", res.GetResult())
	}
}

// func main() {
// 	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Did not connect: %v", err)
// 	}
// 	defer conn.Close()

// 	client := service.NewLongTimeRequestServiceClient(conn)

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	stream, err := client.LongTimeRequestStream(ctx)
// 	if err != nil {
// 		log.Fatalf("Error on get stream: %v", err)
// 	}

// 	req := service.Request{Data: "Test data"}
// 	if err := stream.Send(&req); err != nil {
// 		log.Fatalf("Error while sending stream: %v", err)
// 	}

// 	res, err := stream.Recv()
// 	if err != nil {
// 		log.Fatalf("Error while receiving stream: %v", err)
// 	}

// 	log.Printf("Response from server: %s", res.GetResult())
// }
