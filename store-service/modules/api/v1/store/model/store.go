package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SocialMedia struct {
	PlatformName string `json:"platform_name" bson:"platform_name"`
	Url          string `json:"url" bson:"url"`
}

type Store struct {
	Id                  primitive.ObjectID `json:"id" bson:"_id"`
	Name                string             `json:"name" bson:"name"`
	Description         string             `json:"description,omitempty" bson:"description,omitempty"`
	ProfileImage        string             `json:"profile_image" bson:"profile_image"`
	SocialMediaAccounts []SocialMedia      `json:"social_media,omitempty" bson:"social_media,omitempty"`
	Slug                string             `json:"slug" bson:"slug"`
	CreatedAt           time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedBy           string             `json:"created_by" bson:"created_by"`
	UpdatedBy           string             `json:"updated_by" bson:"updated_by"`
	IsActive            bool               `json:"is_active" bson:"is_active"`
}
