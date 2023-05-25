package plugin

import (
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/hermanschaaf/cq-source-xkcd/client"
	"github.com/hermanschaaf/cq-source-xkcd/resources"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"xkcd",
		Version,
		schema.Tables{
			resources.ComicsTable(),
		},
		client.New,
	)
}
