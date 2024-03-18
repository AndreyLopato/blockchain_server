package repository

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Client
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) Set(height int, blockchain string) error {

	bcColl := repo.db.Database("blockchain_db").Collection("blockchain")

	_, err := bcColl.InsertOne(context.TODO(), bson.D{{"height", height}, {"bc", blockchain}})
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Get(height int) (error, string) {

	bcColl := repo.db.Database("blockchain_db").Collection("blockchain")

	var result bson.M
	err := bcColl.FindOne(context.TODO(), bson.D{{"height", height}}).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return err, ""
		} else {
			return nil, ""
		}
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return err, ""
	}
	return nil, string(jsonData)
}
