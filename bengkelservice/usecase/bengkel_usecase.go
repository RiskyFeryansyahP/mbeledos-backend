package usecase

import (
	"context"

	"github.com/confus1on/mbeledos/bengkelservice"
	"github.com/confus1on/mbeledos/ent"
)

type BengkelUsecase struct {
	RepoBengkel bengkelservice.RepositoryBengkel
}

func NewBengkelUsecase(repobengkel bengkelservice.RepositoryBengkel) bengkelservice.UsecaseBengkel {
	return &BengkelUsecase{
		RepoBengkel: repobengkel,
	}
}

func (bu *BengkelUsecase) GetAllBengkel(ctx context.Context) ([]*ent.Bengkel, error) {
	result, err := bu.RepoBengkel.ShowAllBengkel(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (bu *BengkelUsecase) SpecificationBengkel(ctx context.Context, kode_bengkel string) (*ent.Bengkel, error) {
	result, err := bu.RepoBengkel.GetSpecificationBengkel(ctx, kode_bengkel)

	if err != nil {
		return nil, err
	}

	return result, nil
}
