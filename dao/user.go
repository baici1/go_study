package dao

type User struct {
	Id       int  `json:"id" valid:"required"`
	Name     string `json:"name" valid:"required"`
	Password string `json:"password" valid:"required"`
}