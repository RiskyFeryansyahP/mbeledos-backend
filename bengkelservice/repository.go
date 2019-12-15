package bengkelservice

import (
	"context"

	"github.com/confus1on/mbeledos/ent"
)

type RepositoryBengkel interface {
	ShowAllBengkel(ctx context.Context) ([]*ent.Bengkel, error)
	GetSpecificationBengkel(ctx context.Context, kode_bengkel string) (*ent.Bengkel, error)
}
