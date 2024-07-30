package entity

type UID struct {
	UID uint32 `json:"user_id" validate:"number,required"`
}

type Passport struct {
	PasspotNumber string `json:"passport" validate:"number,required"`
}

type UserInfo struct {
	Surname    string `json:"surname" validate:"alpha"`
	Name       string `json:"name" validate:"alpha"`
	Patronymic string `json:"patronymic" validate:"alpha"`
	Adress     string `json:"adress"`
}

type User struct {
	Passport Passport
	Info     UserInfo
}
