package entity

type User struct {
	PasspotNumber string `validate:"number,required"`
	Surname       string `validate:"alpha,required" json:"surname"`
	Name          string `validate:"alpha,required" json:"name"`
	Patronymic    string `validate:"alpha,required" json:"patronymic"`
	Adress        string `validate:"required" json:"adress"`
}

var (
	UID string
)
