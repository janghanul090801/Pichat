package user

import (
	"Pichat/pkg/ent"
	"context"
)

type Repository interface{}

type repository struct {
	DBConn  *ent.Client
	Context context.Context
}

func NewRepo(dbconn *ent.Client, ctx context.Context) Repository {
	return &repository{
		DBConn:  dbconn,
		Context: ctx,
	}
}
