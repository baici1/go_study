package dao

type User struct {
	Id       int  `json:"id" valid:"required" form:"id"`
	Name     string `json:"name" valid:"required" form:"name"`
	Password string `json:"password" valid:"required" form:"password" `
}