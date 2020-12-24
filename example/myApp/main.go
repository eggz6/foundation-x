package main

import (
	"context"
	"fmt"

	_ "github.com/guanwenbogit/foundation-x/example/pluginC"

	"github.com/guanwenbogit/foundation-x/example/pluginA/app"

	"github.com/guanwenbogit/foundation-x/plugin"
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
