package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name"`
	Last_Name       *string            `json:"last_name"`
	Password        *string            `json:"password"`
	Email           *string            `json:"email"`
	Phone           *string            `json:"phone"`
	Token           *string            ``
	Refresh_Token   *string
	Created_At      time.Time
	Updated_At      time.Time
	User_ID         string
	UserCart        []ProductUser
	Address_Details []Address
	Order_Status    []Order
}

type Product struct {
	Product_ID   primitive.ObjectID
	Product_Name *string
	Price        *uint64
	Rating       *uint8
	Image        *string
}
type ProductUser struct {
	Product_ID   primitive.ObjectID
	Product_Name *string
	Price        int
	RatingImage  *uint8
	Image        *string
}

type Address struct {
	Address_ID primitive.ObjectID
	House      *string
	Street     *string
	City       *string
	PinCode    *string
}

type Order struct {
	Order_ID       primitive.ObjectID
	Order_Cart     []ProductUser
	Ordered_At     time.Time
	Price          int
	Discount       *int
	Payment_Method Payment
}

type Payment struct {
	Digital bool
	COD     bool
}
