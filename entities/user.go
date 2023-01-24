package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FullName    string             `json:"fullName" bson:"fullName"`
	Email       string             `json:"email" bson:"email"`
	Mobile      string             `json:"mobile" bson:"mobile"`
	Password    string             `json:"password,omitempty" bson:"password"`
	UserName    string             `json:"userName" bson:"userName"`
	IsVerified  bool               `json:"isVerified" bson:"isVerified"`
	IsBlocked   bool               `json:"isBlocked" bson:"isBlocked"`
	IsDeleted   bool               `json:"isDeleted" bson:"isDeleted"`
	CreatedDate time.Time          `json:"createdDate" bson:"createdDate"`
}

func (u User) IsEmpty() bool {
	if *new(User) == u {
		return true
	} else {
		return false
	}
}
