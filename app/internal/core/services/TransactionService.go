package services

import (
	"errors"
	"rinha/app/internal/core/domain"
	"rinha/app/internal/core/ports"
)

type service struct {
	repository ports.AccountsRepository
}

func New() *service {
	return &service{}
}

func (srv *service) Save (transaction domain.Transaction) error {
	// validate transaction fields
	// check validation errors
	return srv.repository.Save(transaction)
}

func (srv *service) Get (id uint) (domain.AccountDetails, error) {
	// validate id exists ( ???? )
	// check validation errors, if it makes sense...
	details, err := srv.repository.Get(id)
	if err != nil {
		return domain.AccountDetails{}, errors.New("details not found")
	}

	return details, nil
}