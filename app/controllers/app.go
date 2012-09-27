package controllers

import (
	"github.com/robfig/revel"
)

type Application struct {
	*rev.Controller
}

func (c Application) Index() rev.Result {
	message := "hello, Revel"
	return c.Render(message)
}
