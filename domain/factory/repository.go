package factory

import "github.com/netoudi/codepay-process-transaction/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
