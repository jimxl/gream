package web

import (
	"os"
	"reflect"
	"regexp"

	"gbs/gream/logger"
	"gbs/gream/web/http_router"

	"github.com/olekukonko/tablewriter"
)

var controllers = map[string]reflect.Type{}
var controllerScopeRe = regexp.MustCompile("web/controllers?(.*)$")

type RInfo struct {
	Prefix, Verb, URI, ControllerAndAction string
}

var urls = []*RInfo{}

func AddRInfo(r *RInfo) {
	urls = append(urls, r)
}

func printUrls() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Verb", "URI", "Controller#Action"})
	for _, rinfo := range urls {
		table.Append([]string{rinfo.Verb, rinfo.URI, rinfo.ControllerAndAction})
	}
	table.SetAutoFormatHeaders(false)
	table.Render()
}

func Register(controller Controller) {
	controllerType := reflect.TypeOf(controller)
	controllerScope := controllerScopeRe.FindStringSubmatch(controllerType.Elem().PkgPath())[1]
	controllers[controllerScope+"/"+controllerType.Elem().Name()] = controllerType
}

func GetController(name string) reflect.Type {
	return controllers[name]
}

func Run() {
	http_router.Run()
}

func Debug() {
	// printControllers()
	printUrls()
}

func printControllers() {
	for name, controller := range controllers {
		logger.Debugf("%v => %v", name, controller)
	}
}
