package plugin

import "context"

// RunPluginByName call a plugin Run with the plugin name
func RunPluginByName(ctx context.Context, name string) error {
	return def.RunPluginByName(ctx, name)
}

// Run run the plugins by names one by one, if there is no names, it run all plugins
func Run(ctx context.Context, names ...string) error {
	return def.Run(ctx, names...)
}

// Perpare call all plugins perpare func
func Perpare(ctx context.Context) error {
	return def.Perpare(ctx)
}

// Register register plugins
func Register(ps ...Pluggable) {
	def.Register(ps...)
}
