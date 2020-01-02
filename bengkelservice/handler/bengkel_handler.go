package handler

import (
	"encoding/json"
	"log"

	"github.com/confus1on/mbeledos/bengkelservice"
	"github.com/confus1on/mbeledos/ent"
	"github.com/valyala/fasthttp"
)

type BengkelHandler struct {
	BengkelUsecase bengkelservice.UsecaseBengkel
}

func NewBengkelHandler(bengkelUsecase bengkelservice.UsecaseBengkel) *BengkelHandler {
	return &BengkelHandler{BengkelUsecase: bengkelUsecase}
}

func (bh *BengkelHandler) GetAllDataBengkel(ctx *fasthttp.RequestCtx) {
	res, err := bh.BengkelUsecase.GetAllBengkel(ctx)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(ctx).Encode(res)
}

func (bh *BengkelHandler) GetDataBengkel(ctx *fasthttp.RequestCtx) {
	kode_bengkel := ctx.UserValue("kode_bengkel").(string)

	result, err := bh.BengkelUsecase.SpecificationBengkel(ctx, kode_bengkel)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(ctx).Encode(result)
}

func (bh *BengkelHandler) GetNearestDataBengkel(ctx *fasthttp.RequestCtx) {
	var bengkel ent.Bengkel

	body := ctx.Request.Body()

	json.Unmarshal(body, &bengkel)

	result, err := bh.BengkelUsecase.GetNearestBengkel(ctx, bengkel.Latitude, bengkel.Longitude)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(ctx).Encode(result)
}
