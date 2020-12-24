package main

import (
	"context"
	"fmt"

	"github.com/guanwenbogit/foundation-x/example/pluginA/app"
	"github.com/guanwenbogit/foundation-x/plugin"

	_ "github.com/guanwenbogit/foundation-x/example/pluginB"
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
