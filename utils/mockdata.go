package utils

// import (
// 	"testing"

// 	"github.com/birdglove2/nitad-backend/api/category"
// 	"github.com/birdglove2/nitad-backend/api/project"
// 	"github.com/birdglove2/nitad-backend/api/subcategory"
// 	"github.com/stretchr/testify/require"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func AddMockSubcategory(t *testing.T) *subcategory.Subcategory {
// 	dummySubcate := subcategory.Subcategory{
// 		Title: "dummy subcate title",
// 		Image: "dummy subcate image url",
// 	}

// 	adddedSubcategory, err := subcategory.Add(&dummySubcate)
// 	require.Equal(t, err, nil)
// 	require.Equal(t, dummySubcate.Title, adddedSubcategory.Title)
// 	require.Equal(t, dummySubcate.Image, adddedSubcategory.Image)
// 	require.NotEqual(t, nil, adddedSubcategory.ID)

// 	return adddedSubcategory
// }

// func AddMockCategory(t *testing.T, subcate *subcategory.Subcategory) *category.Category {
// 	dummyCate := category.Category{
// 		Title: "dummy cate title",
// 	}

// 	addedCategory, err := category.Add(&dummyCate, []primitive.ObjectID{subcate.ID})
// 	require.Equal(t, err, nil)
// 	require.Equal(t, dummyCate.Title, addedCategory.Title)
// 	require.Equal(t, dummyCate.Subcategory, addedCategory.Subcategory)
// 	require.NotEqual(t, nil, addedCategory.ID)

// 	return addedCategory
// }

// func AddMockProject(t *testing.T, cate *category.Category) *project.Project {
// 	dummyProj := project.Project{
// 		Title:       "dummy proj title",
// 		Description: "dumym proj description",
// 		Authors:     []string{"dumym proj Authors"},
// 		Emails:      []string{"dumym proj Emails"},
// 		Inspiration: "dumym proj Inspiration",
// 		Abstract:    "dumym proj Abstract",
// 		Images:      []string{"dumym proj Images"},
// 		Videos:      []string{"dumym proj Videos"},
// 		Keywords:    []string{"dumym proj Keywords"},
// 		Report:      "dumym proj Report",
// 		VirtualLink: "dumym proj VirtualLink",
// 		Status:      "dumym proj Status",
// 		Category:    []category.Category{*cate},
// 	}
// 	addedProject, err := project.Add(&dummyProj)
// 	require.Equal(t, err, nil)
// 	require.Equal(t, dummyProj.Title, addedProject.Title)
// 	require.Equal(t, dummyProj.Category, addedProject.Category)
// 	require.NotEqual(t, nil, addedProject.ID)
// 	return addedProject
// }

// func DeleteMock(t *testing.T, proj *project.Project, cate *category.Category, subcate *subcategory.Subcategory) {
// 	err := subcategory.Delete(subcate.ID)
// 	require.Nil(t, err, "Delete subcate failed")

// 	err = category.Delete(cate.ID)
// 	require.Nil(t, err, "Delete cate failed")

// 	err = project.Delete(proj.ID)
// 	require.Nil(t, err, "Delete proj failed")

// }
