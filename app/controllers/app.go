package controllers

import (
	"github.com/revel/revel"
)

// App is struct for context of request & responses
type App struct {
	*revel.Controller
}

// Index method to render index page
func (c App) Index() revel.Result {
	return c.Render()
}

// Articles method to list articles given ofset and limit
func (c App) Articles(offset int, limit int) revel.Result {
	return c.Render(offset)
}
