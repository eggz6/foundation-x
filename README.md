## Foudation-x

### A pluggable framework for developping, it can make your app pluggable

## Get Start

### Import 
```
go get -u -v github.com/guanwenbogit/foudation-x/plugin
```

```
import "github.com/guanwenbogit/foudation-x/plugin"
```

### implement Pluggable in the plugin project, for example:
```
package myPlugin
...
type my struct {
    name string
}

func (m *my) Name() string {
     return my.name
}

func (m *my) Run(ctx context.Context) error {
	log.Print("%s running. ", m.name)

	val := ctx.Value("key")
	  // do something with val
        ...

	return nil
}

func (m *my) Perpare(ctx context.Context) error {
	log.Print("%s perpared. ", m.name)
	val := ctx.Value("key")
	// do something with val
        ...

	return nil
}

```

### Register plugin in the plugin project with init func 
```
improt "github.com/guanwenbogit/foudation-x/plugin"
...
func init() {
    plugin.Register(&my{name: "my-plugin"})
}
...  
```

### Init the plugin in the maintain app, for example: 

```
package app

improt "github.com/guanwenbogit/foudation-x/plugin"
import _ "myPlugin"
...

func init() {
	err := plugin.Perpare(context.TODO())
    if err != nil {
	    panic(fmt.Sprintf("init failed. err=%v", err))
    }
}

func appDo(){
	plugin.Run(context.TODO(),"myPlugin-1","myPlugin-2")
}
```

### 
