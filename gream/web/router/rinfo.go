package router

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type RInfo struct {
	Prefix, Verb, URI, ControllerAndAction string
}

var urls = []*RInfo{}

func AddRInfo(r *RInfo) {
	urls = append(urls, r)
}

func PrintUrls() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Verb", "URI", "Controller#Action"})
	for _, rinfo := range urls {
		table.Append([]string{rinfo.Verb, rinfo.URI, rinfo.ControllerAndAction})
	}
	table.SetAutoFormatHeaders(false)
	table.Render()
}
