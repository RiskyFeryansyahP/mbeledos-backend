package repository

import (
	"context"
	"fmt"

	"github.com/confus1on/mbeledos/ent"
	"github.com/confus1on/mbeledos/ent/transaction"
	"github.com/confus1on/mbeledos/transactionservice"
)

type TransactionRepository struct {
	DB *ent.Client
}

func NewRepositoryTransaction(db *ent.Client) transactionservice.RepositoryTransaction {
	return &TransactionRepository{DB: db}
}

func (tr *TransactionRepository) InsertTransaction(ctx context.Context, transaction ent.Transaction) error {
	_, err := tr.DB.Transaction.
		Create().
		SetOrderphone(transaction.Orderphone).
		SetNamabengkel(transaction.Namabengkel).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("Error when insert transaction : %v" + err.Error())
	}

	return nil
}

func (tr *TransactionRepository) ShowCustomerTransaction(ctx context.Context, orderphone string) ([]*ent.Transaction, error) {
	res, err := tr.DB.Transaction.Query().
		Where(
			transaction.OrderphoneEQ(orderphone),
		).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("Error select transaction %+v ", err)
	}

	return res, nil
}
