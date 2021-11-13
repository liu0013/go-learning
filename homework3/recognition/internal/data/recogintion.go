package data

import (
	"context"
	"recognition/internal/biz"
)

type recogRepo struct {
	data *Data
}

func NewRecogintionRepo(data *Data) biz.GreeterRepo {
	return &recogRepo{
		data: data,
	}
}

func (r *recogRepo) CreateRecogition(ctx context.Context, rg *biz.Greeter) error {
	//TODO
	return nil
}

func (r *recogRepo) GetRecognition(ctx context.Context, rg *biz.Greeter) error {
	//TODO
	return nil
}
