package service

type LoginService interface {
	LoginUser(phone string, password string) bool
}
type loginInformation struct {
	phone    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		phone:    "+964111100",
		password: "12341234",
	}
}
func (info *loginInformation) LoginUser(phone string, password string) bool {
	return info.phone == phone && info.password == password
}
