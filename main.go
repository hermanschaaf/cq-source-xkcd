package main

import (
	"github.com/cloudquery/plugin-sdk/v3/serve"
	"github.com/hermanschaaf/cq-source-xkcd/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
