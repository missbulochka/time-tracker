package entity

type UIDRequest struct {
	UID uint32 `json:"user_id" validate:"number,required"`
}

type User struct {
	PasspotNumber string `json:"passport" validate:"number,required"`
	Surname       string `json:"surname" validate:"alpha"`
	Name          string `json:"name" validate:"alpha"`
	Patronymic    string `json:"patronymic" validate:"alpha"`
	Adress        string `json:"adress"`
}
