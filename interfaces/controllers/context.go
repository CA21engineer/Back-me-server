package controllers

type Context interface {
	Param(string) string
	DefaultQuery(key, defaultValue string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}
