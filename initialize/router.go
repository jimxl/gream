package initialize

import "gbs/web"
import _ "gbs/web/controllers"

func init() {
	router := &web.Router{}
	router.Draw()
}
