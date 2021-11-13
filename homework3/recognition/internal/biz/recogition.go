package biz

import (
	"context"
)

type Recogition struct {
	Comment string
}

type RecongitionRepo interface {
	CreateRecogition(context.Context, *Recogition) error
	GetRecogition(context.Context, *Recogition) error
}

type RecogitionUsecase struct {
	repo RecongitionRepo
}

func NewRecongitionUsecase(repo RecongitionRepo) *RecogitionUsecase {
	return &RecogitionUsecase{repo: repo}
}

func (uc *RecogitionUsecase) Create(ctx context.Context, g *Recogition) error {
	return uc.repo.CreateRecogition(ctx, g)
}

func (uc *RecogitionUsecase) Get(ctx context.Context, g *Recogition) error {
	return uc.repo.CreateRecogition(ctx, g)
}
