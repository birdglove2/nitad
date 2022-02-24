package category

import (
	"github.com/birdglove2/nitad-backend/errors"
	"github.com/birdglove2/nitad-backend/functions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// use for checking ids from ProjectRequest
// receive array of categoryIds, then
// find and return non-duplicated categories, and their ids
// return []Category
func FindByIds(cids []string) ([]CategoryDB, []primitive.ObjectID, errors.CustomError) {
	var objectIds []primitive.ObjectID
	var categories []CategoryDB

	cids = functions.RemoveDuplicateIds(cids)

	for _, cid := range cids {
		oid, err := functions.IsValidObjectId(cid)
		if err != nil {
			return categories, objectIds, err
		}

		categoryDB, err := GetByIdCleanup(oid)
		if err != nil {
			return categories, objectIds, err
		}

		objectIds = append(objectIds, oid)
		categories = append(categories, categoryDB)

	}

	return categories, objectIds, nil
}

// validate requested string of a single categoryId
// and return valid objectId, otherwise error
func ValidateId(cid string) (Category, errors.CustomError) {
	var c Category
	objectId, err := functions.IsValidObjectId(cid)
	if err != nil {
		return c, err
	}

	if c, err = GetById(objectId); err != nil {
		return c, err
	}

	return c, nil
}

// convert bson to category
func BsonToCategory(b bson.M) Category {
	var s Category
	bsonBytes, _ := bson.Marshal(b)
	bson.Unmarshal(bsonBytes, &s)
	return s
}

// TODO: this function is written in O(n^3), should find a better way to handle this later.
// merge multiple categories with multiple sids
// such that the finalCate will result in
// multiple categories that contain only relevant subcategories
// need to do this because the GetById of category return all subcategories
// that are in the category, so we need to filter some out
func FilterCatesWithSids(categories []CategoryDB, sids []primitive.ObjectID) ([]CategoryDB, errors.CustomError) {
	finalCate := []CategoryDB{}
	for _, cate := range categories {
		// subcateThatIsInCate := []subcategory.Subcategory{}
		sidsThatIsInCate := []primitive.ObjectID{}
		for _, id := range cate.Subcategory {
			for index, sid := range sids {
				if id == sid {
					sidsThatIsInCate = append(sidsThatIsInCate, id)
					sids = remove(sids, index)
					index--
				}
			}
		}
		// for _, subcate := range cate.Subcategory {
		// 	for index, id := range sids {
		// 		if subcate.ID == id {
		// 			// subcateThatIsInCate = append(subcateThatIsInCate, subcate)
		// 			sidsThatIsInCate = append(sidsThatIsInCate, id)
		// 			sids = remove(sids, index)
		// 			index--
		// 		}
		// 	}
		// }
		// cate.Subcategory = subcateThatIsInCate
		cate.Subcategory = sidsThatIsInCate
		// finalCate = append(finalCate, CategoryDB{
		// 	ID:          cate.ID,
		// 	Title:       cate.Title,
		// 	Subcategory: sidsThatIsInCate,
		// })
		finalCate = append(finalCate, cate)
	}

	if len(sids) > 0 {
		return finalCate, errors.NewBadRequestError("conflict: some subcategories are not in any categories")
	}
	return finalCate, nil
}

// remove the value at index in slice unordered
func remove(slice []primitive.ObjectID, i int) []primitive.ObjectID {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
