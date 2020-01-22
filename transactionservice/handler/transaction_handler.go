package handler

import (
	"encoding/json"
	"log"

	"github.com/confus1on/mbeledos/ent"
	"github.com/confus1on/mbeledos/transactionservice"
	"github.com/valyala/fasthttp"
)

type TransactionHanlder struct {
	TransactionUsecase transactionservice.UsecaseTransaction
}

type Message struct {
	Status int `json:"status"`
}

func NewTransactionHandler(tu transactionservice.UsecaseTransaction) *TransactionHanlder {
	return &TransactionHanlder{TransactionUsecase: tu}
}

func (th *TransactionHanlder) InsertData(ctx *fasthttp.RequestCtx) {
	var transaction ent.Transaction

	body := ctx.Request.Body()

	json.Unmarshal(body, &transaction)

	err := th.TransactionUsecase.InsertTransaction(ctx, transaction)

	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(ctx).Encode(&Message{Status: 200})
}
