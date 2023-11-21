package model

type User struct {
	UserID   string `json:"userid" binding:"required"`
	Fullname string `json:"fullname"   binding:"required"`
	Email    string `json:"email" binding:"required" `
	Phone    string `json:"phone"  `
	Address  string `json:"address" binding:"required"`
	Role     string `json:"role"  binding:"required"`
<<<<<<< HEAD
}
type Task struct {
	TaskID   string `json:"taskid" binding:"required"`
	Taskname string `json:"taskname"   binding:"required"`
	Starting string `json:"starting" binding:"required" `
	Deadline string `json:"deadline"  `
	Catelogy string `json:"catelogy" binding:"required"`
	UserID   string `json:"userid"  binding:"required"`
=======
>>>>>>> 31be63cbb1e09fe917034d1d666f5c6462168e61
}
type Task struct {
	TaskID   string `json:"taskid" binding:"required"`
	Taskname string `json:"taskname"   binding:"required"`
	Starting string `json:"starting" binding:"required" `
	Deadline string `json:"deadline"  `
	Catelogy string `json:"catelogy" binding:"required"`
	UserID   string `json:"userid"  binding:"required"`
}
