package user

type User struct {
	Surname    string `validate:"alpha,required"`
	Name       string `validate:"alpha,required"`
	Patronymic string `validate:"alpha,required"`
	Adress     string `validate:"required"`
}

type PassportNumber struct {
	Passport string `validate:"number,required"`
}

type UserId struct {
	UID string `validate:"numeric,required"`
}
