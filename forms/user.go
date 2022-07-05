package forms

type AddUser struct {
	Name   string `json:"name" binding:"required"`
	Gender string `json:"gender" binding:"required"`
	Age    int    `json:"age" binding:"required"`
}
