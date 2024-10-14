package controllers

type Context interface {
	Peram(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
