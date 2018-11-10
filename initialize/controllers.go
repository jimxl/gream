package initialize

import (
	"gbs/gream/web"

	"gbs/web/controllers"
	"gbs/web/controllers/admin"
)

func init() {
	web.Register(&controllers.HomeController{})
	web.Register(&controllers.UsersController{})
	web.Register(&admin.HomeController{})
}
