package project

import (
	"os"

	"github.com/birdglove2/nitad-backend/api/admin"
	"github.com/birdglove2/nitad-backend/errors"
	"github.com/birdglove2/nitad-backend/gcp"
	"github.com/birdglove2/nitad-backend/redis"
	"github.com/birdglove2/nitad-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func NewController(
	gcpService gcp.Uploader,
	redisService redis.RedisStorage,
	projectRoute fiber.Router,
) {

	controller := &Controller{gcpService, redisService}

	projectRoute.Get("/", controller.ListProject)
	projectRoute.Get("/:projectId", controller.GetProject)

	projectRoute.Use(admin.IsAuth())
	projectRoute.Post("/", AddAndEditProjectValidator, controller.AddProject)
	projectRoute.Put("/:projectId", AddAndEditProjectValidator, controller.EditProject)
	projectRoute.Delete("/:projectId", controller.DeleteProject)

}

type Controller struct {
	gcpService   gcp.Uploader
	redisService redis.RedisStorage
}

// list all projects
func (contc *Controller) ListProject(c *fiber.Ctx) error {
	pq := new(ProjectQuery)

	if err := c.QueryParser(pq); err != nil {
		return err
	}

	projects, paginate, err := FindAll(pq)
	if err != nil {
		return errors.Throw(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "result": projects, "paginate": paginate})
}

// get project by id
func (contc *Controller) GetProject(c *fiber.Ctx) error {
	projectId := c.Params("projectId")

	objectId, err := utils.IsValidObjectId(projectId)
	if err != nil {
		return errors.Throw(c, err)
	}

	if os.Getenv("APP_ENV") != "test" {
		cacheProject := HandleCacheGetProjectById(c, contc.redisService, projectId)

		if cacheProject != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "result": cacheProject})
		}
	}

	result, err := GetById(objectId)
	if err != nil {
		return errors.Throw(c, err)
	}

	IncrementView(objectId, 1)
	if os.Getenv("APP_ENV") != "test" {
		contc.redisService.SetCache(c.Path(), result)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "result": result})
}

// add a project
func (contc *Controller) AddProject(c *fiber.Ctx) error {
	projectBody, ok := c.Locals("projectBody").(*Project)
	if !ok {
		return errors.Throw(c, errors.NewInternalServerError("Add project went wrong!"))
	}

	files, err := utils.ExtractFiles(c, "report")
	if err != nil {
		return errors.Throw(c, err)
	}
	reportURL, err := contc.gcpService.UploadFile(c.Context(), files[0], collectionName)
	if err != nil {
		return errors.Throw(c, err)
	}

	files, err = utils.ExtractFiles(c, "images")
	if err != nil {
		return errors.Throw(c, err)
	}
	imageURLs, err := contc.gcpService.UploadFiles(c.Context(), files, collectionName)
	if err != nil {
		return errors.Throw(c, err)
	}

	projectBody.Report = reportURL
	projectBody.Images = imageURLs

	result, err := Add(projectBody)
	if err != nil {
		// if there is any error, remove the uploaded files from gcp
		imageURLs = append(imageURLs, reportURL)
		contc.gcpService.DeleteFiles(c.Context(), imageURLs, collectionName)
		return errors.Throw(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "result": result})

}

func (contc *Controller) EditProject(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	oid, err := utils.IsValidObjectId(projectId)
	if err != nil {
		return errors.Throw(c, err)
	}

	updateProject, ok := c.Locals("projectBody").(*Project)
	if !ok {
		return errors.Throw(c, errors.NewInternalServerError("Edit project went wrong!"))
	}

	updateProject.ID = oid
	updateProject, err = contc.HandleUpdateReportAndImages(c, updateProject)
	if err != nil {
		return errors.Throw(c, err)
	}

	result, err := Edit(updateProject)
	if err != nil {
		return errors.Throw(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "result": result})
}

// delete the project
func (contc *Controller) DeleteProject(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	objectId, err := utils.IsValidObjectId(projectId)
	if err != nil {
		return errors.Throw(c, err)
	}

	project, err := GetById(objectId)
	if err != nil {
		return err
	}

	contc.gcpService.DeleteFile(c.Context(), project.Report, collectionName)
	contc.gcpService.DeleteFiles(c.Context(), project.Images, collectionName)

	err = Delete(objectId)
	if err != nil {
		return errors.Throw(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "result": "Delete project " + project.Title + " successfully!"})

}
