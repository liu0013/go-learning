package service

import (
	"context"

	v1 "recognition/api/helloworld/v1"
	"recognition/internal/biz"
)

// GreeterService is a greeter service.
type RecognitionService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *RecognitionService) AddRecog(ctx context.Context, comment string) (string, error) {
	//TODO
	return "success", nil
}
