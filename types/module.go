package types

import (
	"fmt"
	"reflect"
)

type Resource struct {
	Name string
	Value interface{}
}

type Moduler interface {
	Requires() (r []string)
	Provides() (r []Resource)
	AfterInstall(m Module) error
}

type PModuler interface {
	Moduler
	Provide(Name string, Value interface{}) error
}

type BaseModuler struct {
	Resources []Resource
	Namespace string
}

func (b *BaseModuler) Requires() (r []string) {
	return nil
}

func (b *BaseModuler) Provides() []Resource {
	return b.Resources
}

func (b *BaseModuler) Provide(Name string, Value interface{}) (err error) {
	for _, res := range b.Resources {
		if res.Name == Name {
			return fmt.Errorf("duplicate resource name %v", Name)
		}
	}
	b.Resources = append(b.Resources, Resource{Name, Value})
	return nil
}

func (b *BaseModuler) Replace(Name string, Value interface{}) (err error) {
	for i, res := range b.Resources {
		if res.Name == Name {
			b.Resources[i].Value = Value
			return nil
		}
	}
	b.Resources = append(b.Resources, Resource{Name, Value})
	return nil
}

func (b BaseModuler) AfterInstall(m Module) error {
	return nil
}

type Module map[string]interface{}

func (m Module) Require(req string) interface{} {
	return m[req]
}

func (m Module) Install(moduler Moduler) (err error) {
	for _, res := range moduler.Provides() {
		if _, ok := m[res.Name]; ok {
			return fmt.Errorf("duplicate resource name %v", res.Name)
		}
		m[res.Name] = res.Value
	}

	for _, req := range moduler.Requires() {
		if _, ok := m[req]; !ok {
			return fmt.Errorf("does not meets the requirement %v", req)
		}
	}

	err = moduler.AfterInstall(m)
	if err != nil {
		return
	}
	return nil
}

func (m Module) Debug(logger Logger) {
	for k, v := range m {
		logger.Debug("module installed", "path", k, "name", reflect.TypeOf(v))
	}
}
