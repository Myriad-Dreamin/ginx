package types

import (
	"fmt"
	"testing"
)

type myModule struct {
	b *int
}

func (m myModule) Requires() (r []string) {
	return nil
}

func (m myModule) Provides() (r []Resource) {
	return []Resource{
		{"mymodule.a", 1},
		{"mymodule.b", m.b},
	}
}

func (m myModule) AfterInstall(Module) error {
	return nil
}

type subModule struct {
	b *int
}

func (m subModule) Requires() (r []string) {
	return []string{"mymodule.a", "mymodule.b"}
}

func (m subModule) Provides() (r []Resource) {
	return nil
}

func (m subModule) AfterInstall(mm Module) error {
	fmt.Println(mm.Require("mymodule.a"), *mm.Require("mymodule.b").(*int))
	return nil
}


func TestModule_Install(t *testing.T) {
	var m = make(Module)
	var mm = &myModule{b: new(int)}
	err := m.Install(mm)
	if err != nil {
		t.Error(err)
		return
	}
	var mmm = &subModule{}
	err = m.Install(mmm)
	if err != nil {
		t.Error(err)
		return
	}
	*mm.b = 1

	fmt.Println(m.Require("mymodule.a"), *m.Require("mymodule.b").(*int))
}

