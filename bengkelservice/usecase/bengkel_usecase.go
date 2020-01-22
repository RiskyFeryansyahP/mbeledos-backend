package usecase

import (
	"context"
	"math"

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

func (bu *BengkelUsecase) GetNearestBengkel(ctx context.Context, latitude, longitude float64) ([]*ent.Bengkel, error) {
	var bengkels []*ent.Bengkel

	result, err := bu.GetAllBengkel(ctx)
	if err != nil {
		return nil, err
	}

	for key, _ := range result {
		distance := math.Acos(
			math.Cos(latitude*math.Pi/180)*
				math.Cos(result[key].Latitude*math.Pi/180)*
				math.Cos((result[key].Longitude*math.Pi/180)-(longitude*math.Pi/180))+
				math.Sin(latitude*math.Pi/180)*
					math.Sin(result[key].Latitude*math.Pi/180),
		) * 6371

		if distance <= 5 {
			bengkel := &ent.Bengkel{
				ID:            result[key].ID,
				KodeBengkel:   result[key].KodeBengkel,
				NamaBengkel:   result[key].NamaBengkel,
				AlamatBengkel: result[key].AlamatBengkel,
				Latitude:      result[key].Latitude,
				Longitude:     result[key].Longitude,
				Phonenumber:   result[key].Phonenumber,
			}

			bengkels = append(bengkels, bengkel)
		}
	}

	return bengkels, nil
}

func (bu *BengkelUsecase) LoginUserBengkel(ctx context.Context, phone string) (*ent.Bengkel, error) {
	result, err := bu.RepoBengkel.SigninBengkelUser(ctx, phone)

	if err != nil {
		return nil, err
	}

	return result, err
}
