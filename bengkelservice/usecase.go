package bengkelservice

import (
	"context"

	"github.com/confus1on/mbeledos/ent"
)

type UsecaseBengkel interface {
	GetAllBengkel(ctx context.Context) ([]*ent.Bengkel, error)
	GetNearestBengkel(ctx context.Context, latitude, longitude float64) ([]*ent.Bengkel, error)
	SpecificationBengkel(ctx context.Context, kode_bengkel string) (*ent.Bengkel, error)
	LoginUserBengkel(ctx context.Context, phone string) (*ent.Bengkel, error)
}
