package controllers

import (
	"gbs/gream/web"
	"gbs/gream/web/controller"
)

func init() {
	web.Register(&UsersController{})
}

type UsersController struct {
	controller.BaseController
}

func (c *UsersController) Index() {
	c.RenderText("users controller index action")
}

func (c *UsersController) New() {
	c.RenderText("users controller new action")
}

func (c *UsersController) Create() {
	c.RenderText("users controller create action")
}

func (c *UsersController) Show() {
	c.RenderText("users controller show action")
}

func (c *UsersController) Edit() {
	c.RenderText("users controller edit action")
}

func (c *UsersController) Update() {
	c.RenderText("users controller update action")
}

func (c *UsersController) Destroy() {
	c.RenderText("users controller destroy action")
}
