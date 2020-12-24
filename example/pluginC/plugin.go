package pluginC

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/eggz6/foundation-x/example"
	"github.com/eggz6/foundation-x/plugin"
)

const PluginC = "plugin-c"

var (
	def *myPlugin
	mux http.ServeMux
)

func init() {
	def = &myPlugin{name: PluginC}
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

	mux.HandleFunc("/pluginC", pluginBHandler)

	return nil
}

func (my *myPlugin) Perpare(ctx context.Context) error {
	log.Print("%s perpared. ", my.name)
	return nil
}

func pluginBHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(PluginC))
}
