package database

import (
	"context"

	"github.com/jakoblorz/autofone/database/mongo"
	"github.com/jakoblorz/autofone/database/sql"
)

type I interface {
	Mongo() *mongo.I
	SQL() *sql.I
}

type i struct {
	mongo *mongo.I
	sql   *sql.I
}

func (i *i) Mongo() *mongo.I {
	return i.mongo
}

func (i *i) SQL() *sql.I {
	return i.sql
}

type option struct {
	fn func(i *i) error
}

func (i *i) applyOption(o *option) error {
	return o.fn(i)
}

func Open(opts ...*option) (I, error) {
	instance := new(i)
	var err error
	for _, o := range opts {
		err = instance.applyOption(o)
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}

func WithMongo(url string, dbName string) *option {
	return &option{
		fn: func(i *i) (err error) {
			if i.mongo == nil {
				i.mongo = new(mongo.I)
			}
			_, err = i.mongo.Connect(url, dbName)
			return
		},
	}
}

func WithGCP_SQL(ctx context.Context, dsn string, bucket string) *option {
	return &option{
		fn: func(i *i) (err error) {
			if i.sql == nil {
				i.sql = new(sql.I)
			}
			_, err = i.sql.GCP(ctx, dsn, bucket)
			return
		},
	}
}
