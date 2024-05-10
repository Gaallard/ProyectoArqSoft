package users

import "backend/dto/users"

func Login(request users.LoginRequest) users.LoginResponse {

	// Validar contra la base de datos

	return users.UsersDto{
		Token: "abcdef123456",
	}
}
