package entity

type IDRequest struct {
	UID   string `json:"user_id" validate:"number,required"`
	Alias string `json:"alias,omitempty"`
}

type User struct {
	PasspotNumber string `validate:"number,required"`
	Surname       string `validate:"alpha,required" json:"surname"`
	Name          string `validate:"alpha,required" json:"name"`
	Patronymic    string `validate:"alpha,required" json:"patronymic"`
	Adress        string `validate:"required" json:"adress"`
}
