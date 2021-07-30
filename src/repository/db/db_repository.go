package db

import (
	"github.com/msd79/bookstore_users-api/utils/errors"
	"github.com/msd79/bookstrore_outh-api/src/domain/access_token"
)

func NewRepository() DBRepository {
	return &dbRepository{}
}

type DBRepository interface {
	GetByID(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r dbRepository) GetByID(string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
