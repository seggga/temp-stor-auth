package grpc

import (
	"context"

	"github.com/seggga/temp-stor-auth/internal/domain/models"
	"github.com/seggga/temp-stor-auth/internal/ports"
)

type Service struct {
	stor ports.UserStorage
}

func New(stor ports.UserStorage) *Service {
	return &Service{
		stor: stor,
	}
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}

func (s *Service) Stop() {}

func (s *Service) Validate(ctx context.Context, tokens *models.Token) error {

	return nil
}
