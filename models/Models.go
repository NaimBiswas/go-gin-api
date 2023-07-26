package Models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id              primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName       string             `json:"firstName,omitempty" validate:"required"`
	LastName        string             `json:"lastName" validate:"required"`
	AcStatus        string             `json:"acStatus,omitempty" validate:"required"`
	Email           string             `json:"email"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"`
	IsActive        bool               `json:"isActive"`
	IsBlocked       bool               `json:"isBlocked"`
	IsDeleted       bool               `json:"isDeleted"`
	IsEmailVerified bool               `json:"isEmailVerified"`
	PaymentDetails  any                `json:"paymentDetails"`
	Role            any                `json:"role"`
}
