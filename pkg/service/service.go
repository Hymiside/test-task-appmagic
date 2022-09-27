package service

import "github.com/Hymiside/test-task-appmagic/pkg/repository"

type Service struct {
	redis *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{redis: r}
}
