package reflectutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Fn struct {
}

func (f *Fn) DoSomething() error {
	return nil 
}

type Obj struct {
}

func (o *Obj) DoSomething() error {
	return nil 
}

type Handler interface {
	DoSomething() error
}

func Test_GetFnNameByReflect(t *testing.T) {
	var fn Fn
	var obj Obj
	res, err := GetFnNameByReflect(fn)
	assert.Nil(t, err)
	t.Log(res)
	res, err = GetFnNameByReflect(&obj)
	assert.Nil(t, err)
	t.Log(res)
}
