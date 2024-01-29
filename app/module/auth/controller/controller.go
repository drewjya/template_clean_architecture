package controller

type Controller struct {
	Auth AuthController
}

func NewController() *Controller {
	return &Controller{}
}
