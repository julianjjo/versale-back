package model

type Customer struct {
	CustomerId string `json:"customer_id" bson:"customer_id"`
	FirstName  string `json:"first_name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	Email      string `json:"email" bson:"email"`
	TypeId     string `json:"type_id" bson:"type_id"`
	DocumentId string `json:"document_id" bson:"document_id"`
	Age        int    `json:"age" bson:"age"`
	Password   string `json:"password" bson:"password"`
}
