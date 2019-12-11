package config

import (
	"context"
	"log"

	"github.com/confus1on/mbeledos/ent"
	_ "github.com/lib/pq"
)

var db *ent.Client

func New() *ent.Client {
	if db != nil {
		return db
	}

	client, err := ent.Open("postgres", "postgres://ztsuiwtp:PQEjOenZ-xM8IyKEbe70uVO4KQy-QP6e@arjuna.db.elephantsql.com:5432/ztsuiwtp")
	if err != nil {
		log.Fatalf("Failed to connect into postgreSQL %v \n", err.Error())
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}
