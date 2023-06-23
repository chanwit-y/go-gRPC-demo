package service

import (
	"io"
	"log"
	"strconv"
	"time"
)

type demoService struct{}

func NewDemoService() demoService {
	return demoService{}
}

func (d demoService) mustEmbedUnimplementedLongTimeRequestServiceServer() {}

func (d demoService) LongTimeRequestStream(stream LongTimeRequestService_LongTimeRequestStreamServer) error {
	// for {
	// 	req, err := stream.Recv()
	// 	if err == io.EOF {
	// 		// The client has closed the stream.
	// 		return nil
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("Error while reading client stream: %v", err)
	// 	}

	// 	// Handle the request and send a response.
	// 	response := Response{Result: "Received: " + req.GetData()}
	// 	if err := stream.Send(&response); err != nil {
	// 		log.Fatalf("Error while sending response to client: %v", err)
	// 	}
	// }

	for {
		_, err := stream.Recv()
		if err == io.EOF {
			// The client has closed the stream.
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		// Handle the request and send a response.
		for i := 0; i < 10; i++ {
			res := &Response{Result: "Message " + strconv.Itoa(i)}
			if err := stream.Send(res); err != nil {
				return err
			}
			time.Sleep(1 * time.Second)
		}
		return nil
	}

}

func (d demoService) LongTimeRequestStream2(stream LongTimeRequestService_LongTimeRequestStream2Server) error {
	return nil
}
