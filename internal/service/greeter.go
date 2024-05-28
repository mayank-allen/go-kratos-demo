package service

import (
	"context"
	"demo/internal/data/models"
	"fmt"

	v1 "demo/api/helloworld/v1"
	"demo/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &models.Greeter{M1: in.B1, M2: in.B2})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + fmt.Sprint(*g)}, nil
}
