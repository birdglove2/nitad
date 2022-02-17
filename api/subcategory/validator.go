package subcategory

import (
	"github.com/birdglove2/nitad-backend/api/validators"
	"github.com/birdglove2/nitad-backend/errors"
	"github.com/gofiber/fiber/v2"
)

func AddSubcategoryValidator(c *fiber.Ctx) error {
	p := new(SubcategoryRequest)

	if err := c.BodyParser(p); err != nil {
		return errors.Throw(c, errors.NewBadRequestError(err.Error()))
	}

	err := validators.ValidateStruct(*p)
	if err != nil {
		return errors.Throw(c, err)
	}

	return c.Next()
}
