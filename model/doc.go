package model

// Doc for database
type Doc struct {
	ID       string `json:"id" bson:"id"`
	TypeCode string `json:"typecode" bson:"typecode"`
	Comment  string `json:"comment" bson:"comment"`
}
