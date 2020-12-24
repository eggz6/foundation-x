package main

import (
	"context"
	"fmt"

	_ "github.com/eggz6/foundation-x/example/pluginC"

	"github.com/eggz6/foundation-x/example/pluginA/app"

	"github.com/eggz6/foundation-x/plugin"
)

func init() {
	err := plugin.Perpare(context.TODO())
	if err != nil {
		panic(fmt.Sprintf("init failed. err=%v", err))
	}
}

func main() {
	app.Launch()
}
