package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *Mongo) getItemsCollection() *mongo.Collection {
	if m.collections == nil {
		m.collections = make(map[string]*mongo.Collection)
	}
	if _, ok := m.collections["items"]; !ok {
		m.collections["items"] = m.DB.Collection("items")
	}
	return m.collections["items"]
}

type Item struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Price *float32 `json:"price"`
}

func (m *Mongo) ListItems(ctx context.Context) ([]*Item, error) {
	return m.getItemsByFilter(ctx, bson.M{})
}

func (m *Mongo) GetItemById(ctx context.Context, id string) (*Item, error) {
	return m.getitemByFilter(ctx, bson.M{"id": id})
}

func (m *Mongo) SaveItem(ctx context.Context, item *Item) error {
	collection := m.getItemsCollection()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "id", Value: item.ID}}
	update := bson.D{{Key: "$set", Value: item}}

	_, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) DeleteItem(ctx context.Context, id string) error {
	collection := m.getItemsCollection()

	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongo) getItemsByFilter(ctx context.Context, filter interface{}) ([]*Item, error) {
	collection := m.getItemsCollection()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var items []*Item
	for cursor.Next(ctx) {
		var item Item
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}

func (m *Mongo) getitemByFilter(ctx context.Context, filter interface{}) (*Item, error) {
	// Get a handle to the users collection
	collection := m.getItemsCollection()

	// Execute the find query
	var item Item
	err := collection.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return &item, nil
}
