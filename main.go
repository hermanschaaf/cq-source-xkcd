package main

import (
	"github.com/cloudquery/plugin-sdk/serve"
	"github.com/hermanschaaf/cq-source-xkcd/plugin"
)

func main() {
	serve.Source(plugin.Plugin())
}
