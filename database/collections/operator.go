package collections

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type operator struct {
	*mongo.Collection
}

func (o *operator) InsertOne(ctx context.Context, model interface{}) (interface{}, error) {
	res, err := o.Collection.InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (o *operator) FindOne(ctx context.Context, filter interface{}, target interface{}) error {
	return o.Collection.FindOne(ctx, filter).Decode(target)
}

func (o *operator) Find(ctx context.Context, filter interface{}, target interface{}) error {
	crs, err := o.Collection.Find(ctx, filter)
	if err != nil {
		return err
	}

	return crs.All(ctx, target)
}

func (o *operator) UpdateOne(ctx context.Context, filter interface{}, target interface{}) error {
	res, err := o.Collection.UpdateOne(ctx, filter, target)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 && res.ModifiedCount == 0 && res.UpsertedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
