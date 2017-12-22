package models

//Post is struct for save data post
type Notes struct {
	ID          uint   `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
}
