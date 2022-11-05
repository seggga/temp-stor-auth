package auth

import (
	"context"

	"github.com/seggga/temp-stor-auth/internal/domain/models"
	"github.com/seggga/temp-stor-auth/internal/ports"
)

// Service implements main auth logic
type Service struct {
	db ports.UserStorage
}

// New creates a new auth service
func New(db ports.UserStorage) *Service {
	return &Service{
		db: db,
	}
}

// Validate checks provided password
func (s *Service) Validate(ctx context.Context, token models.Token) (string, error) {
	// that is stub
	return "", nil
}

// Login produces token
func (s *Service) Login(ctx context.Context, user, password string) (models.Token, error) {
	// that is stub
	var token models.Token
	return token, nil
}
