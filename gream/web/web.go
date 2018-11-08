package web

import (
	"os"

	"gbs/gream/web/http_router"

	"github.com/olekukonko/tablewriter"
)

// TODO: 移动到子模块中去
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

func Run() {
	http_router.Run()
}

func Debug() {
	// PrintControllers()
	printUrls()
}
