package service

import (
	"context"
	"golang-keycloak/dto"
	"golang-keycloak/repository"
)

type (
	TodoSerivce interface {
		Fatch(ctx context.Context) (*[]dto.ResTodo, error)
	}

	todoImpl struct {
		repository repository.TodoRepository
	}
)

func NewTodoService(repository repository.TodoRepository) TodoSerivce {
	return todoImpl{repository: repository}
}

func (s todoImpl) Fatch(ctx context.Context) (*[]dto.ResTodo, error) {
	req, err := s.repository.Fatch(ctx)
	if err != nil {
		return nil, err
	}

	var todo []dto.ResTodo
	for _, row := range req {
		todo = append(todo, dto.ResTodo{Name: row.Name})
	}

	return &todo, nil
}
