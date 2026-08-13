package service

import (
	"context"

	"example.com/security/internal/repository"
)

type JsonService[T any] struct {
	Store repository.JsonRepository[T]
}

func (s *JsonService[T]) Load(ctx context.Context, v string) (*T, error) {
	return s.Store.Load(ctx, v)
}

func (s *JsonService[T]) Save(ctx context.Context, v string, d T) error {
	return s.Store.Save(ctx, v, &d)
}
