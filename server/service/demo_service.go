package service

import (
	"io"
	"log"
)

type demoService struct{}

func NewDemoService() demoService {
	return demoService{}
}

func (d demoService) mustEmbedUnimplementedLongTimeRequestServiceServer() {}

func (d demoService) LongTimeRequestStream(stream LongTimeRequestService_LongTimeRequestStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// The client has closed the stream.
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		// Handle the request and send a response.
		response := Response{Result: "Received: " + req.GetData()}
		if err := stream.Send(&response); err != nil {
			log.Fatalf("Error while sending response to client: %v", err)
		}
	}
}