package repository

import (
	"context"
	"golang-keycloak/model"
)

type (
	TodoRepository interface {
		Fatch(ctx context.Context) ([]model.Todo, error)
	}

	todoImpl struct {
	}
)

func NewTodoRepostiroy() TodoRepository {
	return todoImpl{}
}

func (r todoImpl) Fatch(ctx context.Context) ([]model.Todo, error) {

	return []model.Todo{
		{Name: "Helo"},
		{Name: "Heho"},
	}, nil
}
