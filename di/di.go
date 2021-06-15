package di

import (
	"fmt"
	"reflect"
)

type Container struct {
	providers map[reflect.Type]provider

	results map[reflect.Type]reflect.Value
}

type provider struct {
	value  reflect.Value
	params []reflect.Type
}

func New() *Container {
	return &Container{
		providers: map[reflect.Type]provider{},
		results:   map[reflect.Type]reflect.Value{},
	}
}

func isError(t reflect.Type) bool {
	if t.Kind() != reflect.Interface {
		return false
	}
	return t.Implements(reflect.TypeOf((*error)(nil)).Elem())
}







func (c *Container) Provide(constructor interface{}) error {
	v := reflect.ValueOf(constructor)
	if v.Kind() != reflect.Func {
		return fmt.Errorf("constructor must be a func ")
	}
	vt := v.Type()
	params := make([]reflect.Type, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		params[i] = vt.In(i)
	}

	results := make([]reflect.Type, vt.NumIn())
	for i := 0; i < vt.NumIn(); i++ {
		results[i] = vt.Out(i)/**/
	}

	provider := provider{
		value:  v,
		params: params,
	}

	for _, result := range results {
		if isError(result) {
			continue
		}
		if _, ok := c.providers[result]; ok {
			return fmt.Errorf("%s had a provider", result)
		}
		c.providers[result] = provider

	}
	return nil
}

