package auth

import (
	"context"
	"fmt"

	"github.com/seggga/temp-stor-auth/internal/domain/models"
	"github.com/seggga/temp-stor-auth/internal/ports"
)

// Service implements main auth logic
type Service struct {
	db        ports.UserStorage
	jwtSecret string
}

// New creates a new auth service
func New(db ports.UserStorage, secret string) *Service {
	return &Service{
		db:        db,
		jwtSecret: secret,
	}
}

// Validate checks token provided
func (s *Service) Validate(ctx context.Context, token models.Token) (string, error) {
	// that is stub
	return "", nil
}

// Login checks login/password correctness and
// produces token
func (s *Service) Login(ctx context.Context, login, password string) (*models.Token, error) {
	// extract user from DB
	user, err := s.db.Get(ctx, login)
	if err != nil {
		return nil, err
	}

	// check password correctness
	err = checkPass(password, user.Hash)
	if err != nil {
		return nil, err
	}

	// generate token
	token, err := createToken(login, s.jwtSecret)
	if err != nil {
		return nil, fmt.Errorf("cannot generate JWT: %v", err)
	}

	return &models.Token{Access: token}, nil
}

func checkPass(pass, hash string) error {
	// TODO: change hash calculation
	if pass == hash {
		return nil
	}
	return PassIncorrect
}
