package controller

import "gbs/gream/web/http_router"

type ActionI = func(context http_router.Context)
