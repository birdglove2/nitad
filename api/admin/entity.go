package admin

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty`
	Username string             `bson:"username,omitempty`
	Password []string           `bson:"password,omitempty`
}
