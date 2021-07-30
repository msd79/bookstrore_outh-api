package access_token

import (
	"github.com/msd79/bookstrore_outh-api/src/repository/db"
	"github.com/msd79/bookstrore_outh-api/src/utils/errors"
)

type Repository interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetByID(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo db.DBRepository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetByID(string) (*AccessToken, *errors.RestErr) {
	return nil, nil
}
