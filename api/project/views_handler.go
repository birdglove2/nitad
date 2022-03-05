package project

import (
	"github.com/birdglove2/nitad-backend/database"
	"github.com/birdglove2/nitad-backend/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func IncrementView(id primitive.ObjectID, val int) {
	projectCollection, ctx := database.GetCollection(collectionName)

	_, err := projectCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{Key: "$inc", Value: bson.D{{Key: "views", Value: val}}},
		},
	)

	if err != nil {
		zap.S().Warn("Incrementing view error: ", err.Error())
	}
}

func IncrementViewCache(r redis.RedisStorage, id string, views int) {
	key := "views" + id
	// log.Println("9")
	countInt := r.GetCacheInt(key)
	// log.Println("10")
	if countInt != 0 {
		r.SetCacheInt(key, countInt+1)
		zap.S().Info("incrementing view of ", key, " = ", countInt+1)
		return
	}
	r.SetCacheInt(key, 1)
}
