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
type IResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}
type UserColumnsForExport struct {
	FirstName       string             `json:"firstName" pdfFiled:"First Name"`
	LastName        string             `json:"lastName" pdfFiled:"Last Name"`
	Email           string             `json:"email" pdfFiled:"Email"`
	Role            string             `json:"role" pdfFiled:"Role"`
	UpdatedAt       primitive.DateTime `json:"updatedAt" pdfFiled:"Last Updated"`
	AcStatus        string             `json:"acStatus" pdfFiled:"Account Status"`
	IsActive        bool               `json:"isActive" pdfFiled:"Is Active"`
	IsEmailVerified bool               `json:"isEmailVerified" pdfFiled:"Is Email Verified"`
}

type ProductModel struct {
	ID                 primitive.ObjectID `json:"_id" bson:"_id"`
	Name               string             `json:"name"`
	DiscountPrice      int                `json:"discountPrice"`
	ProductDescription string             `json:"productDescription"`
	ProductImage       string             `json:"productImage"`
	MetaTitle          string             `json:"metaTitle"`
	MetaKeys           string             `json:"metaKeys"`
	MetaDescription    string             `json:"metaDescription"`
	OriginPrice        int                `json:"originPrice"`
	ManufacturerName   string             `json:"manufacturerName"`
	ManufacturerBrand  string             `json:"manufacturerBrand"`
	Price              int                `json:"price"`
	Discount           int                `json:"discount"`
	Category           struct {
		CategoryID string `json:"categoryId"`
		Name       string `json:"name"`
	} `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	IsActive bool `json:"isActive"`
}

type ImportedProductModel struct {
}
