package registration

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterReplyType int8
type RegisterCommandType int8

const (
	SaveProfile RegisterCommandType = iota
	RollbackProfile
	RollbackAuth
	SaveGraph
	UnknownCommand
)

const (
	ProfileSaved RegisterReplyType = iota
	ProfileNotSaved
	GraphSaved
	GraphNotSaved
	AuthSaved
	AuthNotSaved
	UnknownReply
)

type User struct {
	FullName        string             `validate:"omitempty,fullName" json:"fullName,omitempty"    bson:"fullName,omitempty"`
	Country         string             `validate:"omitempty,len=2"  json:"country,omitempty"     bson:"country,omitempty"`
	City            string             `validate:"omitempty,min=3"  json:"city,omitempty"     bson:"city,omitempty"`
	CompanyName     string             `validate:"omitempty,company"  json:"companyName,omitempty" bson:"companyName,omitempty"`
	Email           string             `validate:"omitempty,email"    json:"email,omitempty"       bson:"email,omitempty"`
	Website         string             `validate:"omitempty,url"      json:"website,omitempty"     bson:"website,omitempty"`
	Username        string             `validate:"omitempty,username" json:"username,omitempty"    bson:"username,omitempty"`
	Sex             string             `validate:"omitempty,sex"      json:"sex,omitempty"         bson:"sex,omitempty"`
	Role            string             `json:"role,omitempty" bson:"role,omitempty"`
	Age             int                `validate:"omitempty,gt=12"    json:"age,omitempty"         bson:"age,omitempty"`
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IsPrivate       bool               `json:"isPrivate" bson:"isPrivate"`
	ProfileImageUrl string             `json:"profileImageUrl"  bson:"profileImageUrl"`
}

type RegisterUserCommand struct {
	User User
	Type RegisterCommandType
}

type RegisterUserReply struct {
	User User
	Type RegisterReplyType
}
