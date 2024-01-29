package controller

import "template_clean_architecture/app/module/auth/service"

type authController struct {
	authService service.AuthService
}

type AuthController interface {
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}
