package usecase

import (
	"context"

	"github.com/confus1on/mbeledos/ent"
	"github.com/confus1on/mbeledos/transactionservice"
)

type TransactionUsecase struct {
	RepoTransaction transactionservice.RepositoryTransaction
}

func NewUsecaseTransaction(rt transactionservice.RepositoryTransaction) transactionservice.UsecaseTransaction {
	return &TransactionUsecase{RepoTransaction: rt}
}

func (tu *TransactionUsecase) InsertTransaction(ctx context.Context, transaction ent.Transaction) error {

	err := tu.RepoTransaction.InsertTransaction(ctx, transaction)

	if err != nil {
		return err
	}

	return nil
}

func (tu *TransactionUsecase) SelectCustomerTransaction(ctx context.Context, orderphone string) ([]*ent.Transaction, error) {
	res, err := tu.RepoTransaction.ShowCustomerTransaction(ctx, orderphone)
	if err != nil {
		return nil, err
	}

	return res, nil
}
