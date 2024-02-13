package ports

import "rinha/app/internal/core/domain"

type AccountsRepository interface {
	Get (id uint) (domain.AccountDetails, error)
	Save (transaction domain.Transaction) error
}