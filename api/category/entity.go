package category

import (
	"time"

	"github.com/birdglove2/nitad-backend/api/subcategory"
	"github.com/birdglove2/nitad-backend/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionName = database.COLLECTIONS["CATEGORY"]

type Category struct {
	ID          primitive.ObjectID        `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string                    `bson:"title" json:"title"`
	Subcategory []subcategory.Subcategory `bson:"subcategory" json:"subcategory"`
	CreatedAt   time.Time                 `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time                 `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type CategoryDTO struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string               `bson:"title" json:"title"`
	Subcategory []primitive.ObjectID `bson:"subcategory" json:"subcategory"`
	CreatedAt   time.Time            `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time            `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type CategoryRequest struct {
	ID          primitive.ObjectID `form:"-"`
	Title       string             `form:"title" validate:"required"`
	Subcategory []string           `form:"subcategory"`
	CreatedAt   time.Time          `form:"-"`
	UpdatedAt   time.Time          `form:"-"`
}

type CategorySearch struct {
	ID          primitive.ObjectID              `bson:"_id" json:"id"`
	Title       string                          `bson:"title" json:"title"`
	Subcategory []subcategory.SubcategorySearch `bson:"subcategory" json:"subcategory"`
}
