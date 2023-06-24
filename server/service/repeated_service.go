package service

import (
	context "context"
	"fmt"
)

// GetBeer(context.Context, *RequestRepeated) (*ResponseRepeated, error)
// mustEmbedUnimplementedRepeatedServiceServer()

type repeatedService struct{}

func NewRepeatedService() repeatedService {
	return repeatedService{}
}

func (repeatedService) mustEmbedUnimplementedRepeatedServiceServer() {}

func (repeatedService) GetBeer(ctx context.Context, req *RequestRepeated) (*ResponseRepeated, error) {
	fmt.Printf("Get Name: %v\n", req.GetName())
	res := &ResponseRepeated{
		Requests: []*RequestRepeated{
			req,
		},
	}
	return res, nil
}
