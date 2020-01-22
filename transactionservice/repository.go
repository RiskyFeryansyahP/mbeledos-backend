package transactionservice

import (
	"context"

	"github.com/confus1on/mbeledos/ent"
)

type RepositoryTransaction interface {
	InsertTransaction(ctx context.Context, transaction ent.Transaction) error
}
