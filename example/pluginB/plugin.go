package pluginB

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/guanwenbogit/foundation-x/example"
	"github.com/guanwenbogit/foundation-x/plugin"
)

const PluginB = "plugin-b"

var (
	def *myPlugin
	mux http.ServeMux
)

func init() {
	def = &myPlugin{name: PluginB}
	plugin.Register(def)
}

type myPlugin struct {
	name string
}

func (my *myPlugin) Name() string {
	return my.name
}

func (my *myPlugin) Run(ctx context.Context) error {
	log.Print("%s running. ", my.name)

	val := ctx.Value(example.HttpMux)
	mux, ok := val.(*http.ServeMux)

	if !ok {
		return fmt.Errorf("%s need a *http.ServeMux")
	}

	mux.HandleFunc("/pluginB", pluginBHandler)

	return nil
}

func (my *myPlugin) Perpare(ctx context.Context) error {
	log.Print("%s perpared. ", my.name)
	return nil
}

func pluginBHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(PluginB))
}
