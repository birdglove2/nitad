package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AppendLookupStage(pipe mongo.Pipeline, collectionName string) mongo.Pipeline {
	return append(pipe, bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: collectionName},
		{Key: "localField", Value: collectionName},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: collectionName}}}})
}

func AppendUnsetStage(pipe mongo.Pipeline, field string) mongo.Pipeline {
	return append(pipe, bson.D{{Key: "$unset", Value: field}})
}

func AppendMatchIdStage(pipe mongo.Pipeline, oid primitive.ObjectID) mongo.Pipeline {
	return append(pipe, bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: oid}}}})
}
