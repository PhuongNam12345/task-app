package model

type User struct {
	UserID   string `json:"userid" binding:"required"`
	Fullname string `json:"fullname"   binding:"required"`
	Email    string `json:"email" binding:"required" `
	Phone    string `json:"phone"  `
	Address  string `json:"address" binding:"required"`
	Role     string `json:"role"  binding:"required"`
}
