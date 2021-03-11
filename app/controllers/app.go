package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}
func (c App) Test() revel.Result {
	return c.Render()
}
func (c App) Proceed() revel.Result {
	fname := c.Params.Query.Get("fname")
	lname := c.Params.Query.Get("lname")
	return c.Render(fname, lname)
}
