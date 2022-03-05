package project

import (
	"fmt"
	"os"
	"testing"

	"github.com/birdglove2/nitad-backend/api"
	"github.com/birdglove2/nitad-backend/api/category"
	"github.com/birdglove2/nitad-backend/api/project"
	"github.com/birdglove2/nitad-backend/api/subcategory"
	"github.com/birdglove2/nitad-backend/config"
	"github.com/birdglove2/nitad-backend/database"
	redis "github.com/birdglove2/nitad-backend/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewTestApp(t *testing.T, gcpService *MockUploader, redisService redis.RedisStorage) *fiber.App {
	config.Loadenv()
	database.ConnectDb(os.Getenv("MONGO_URI"))

	app := fiber.New()

	// redisService := redis.Init()
	fmt.Println("check 1")
	app.Use(cache.New(cache.Config{
		Expiration: redis.DefaultCacheExpireTime,
		Storage:    redisService.GetStorage(),
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Path() + "?" + string(c.Request().URI().QueryString())
		},
		Next: func(c *fiber.Ctx) bool {
			isTrue := project.IsGetProjectPath(c) // handle incrementing view in cache
			return isTrue
		},
	}))
	fmt.Println("check 2")

	api.CreateAPI(app, gcpService, redisService)
	fmt.Println("check 3")

	return app
}

func addMockSubcategory(t *testing.T) *subcategory.Subcategory {
	dummySubcate := subcategory.Subcategory{
		Title: "dummy subcate title",
		Image: "dummy subcate image url",
	}

	adddedSubcategory, err := subcategory.Add(&dummySubcate)
	require.Equal(t, err, nil)
	require.Equal(t, dummySubcate.Title, adddedSubcategory.Title)
	require.Equal(t, dummySubcate.Image, adddedSubcategory.Image)
	require.NotEqual(t, nil, adddedSubcategory.ID)

	return adddedSubcategory
}

func addMockCategory(t *testing.T, subcate *subcategory.Subcategory) *category.Category {
	dummyCate := category.Category{
		Title: "dummy cate title",
	}

	addedCategory, err := category.Add(&dummyCate, []primitive.ObjectID{subcate.ID})
	require.Equal(t, err, nil)
	require.Equal(t, dummyCate.Title, addedCategory.Title)
	require.Equal(t, dummyCate.Subcategory, addedCategory.Subcategory)
	require.NotEqual(t, nil, addedCategory.ID)

	return addedCategory
}

func addMockProject(t *testing.T, cate *category.Category) *project.Project {
	dummyProj := project.Project{
		Title:       "dummy proj title",
		Description: "dumym proj description",
		Authors:     []string{"dumym proj Authors"},
		Emails:      []string{"dumym proj Emails"},
		Inspiration: "dumym proj Inspiration",
		Abstract:    "dumym proj Abstract",
		Images:      []string{"dumym proj Images"},
		Videos:      []string{"dumym proj Videos"},
		Keywords:    []string{"dumym proj Keywords"},
		Report:      "dumym proj Report",
		VirtualLink: "dumym proj VirtualLink",
		Status:      "dumym proj Status",
		Category:    []category.Category{*cate},
	}
	addedProject, err := project.Add(&dummyProj)
	require.Equal(t, err, nil)
	require.Equal(t, dummyProj.Title, addedProject.Title)
	require.Equal(t, dummyProj.Category, addedProject.Category)
	require.NotEqual(t, nil, addedProject.ID)
	return addedProject
}

func deleteMock(t *testing.T, proj *project.Project, cate *category.Category, subcate *subcategory.Subcategory) {
	err := subcategory.Delete(subcate.ID)
	require.Nil(t, err, "Delete subcate failed")

	err = category.Delete(cate.ID)
	require.Nil(t, err, "Delete cate failed")

	err = project.Delete(proj.ID)
	require.Nil(t, err, "Delete proj failed")

}
