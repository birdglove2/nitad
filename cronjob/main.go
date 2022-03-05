package cronjob

import (
	"github.com/birdglove2/nitad-backend/api/project"
	"github.com/birdglove2/nitad-backend/redis"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func Init(redisService redis.RedisStorage) {
	c := cron.New()

	cron := &Cron{redisService}
	c.AddFunc("@every 12h", cron.UpdateProjectViews) // every 12 hours
	// c.AddFunc("@every 50s", UpdateProjectViews)      // test

	c.Start()
}

type Cron struct {
	redisService redis.RedisStorage
}

func (cron Cron) UpdateProjectViews() {
	store := cron.redisService.GetStorage()
	for {
		keys, cursor := store.Scan("views")

		for _, key := range keys {
			projectId := key[5:]
			objectId, _ := primitive.ObjectIDFromHex(projectId)

			countInt := cron.redisService.GetCacheInt(key)

			project.IncrementView(objectId, countInt)

			store.Delete(key)
			store.Delete(projectId)
		}

		if cursor == 0 { // no more keys
			break
		}
	}

	zap.S().Info("updating project views")
}
