package middleware

import (
	"sync"
	"testing"
)

func TestFlow(t *testing.T) {
	m := New()
	m.Add(func(c *Context) {
		c.Data["num"] = 0
		t.Logf("a-in: %v\n", c.Data["num"])
		c.Next()
		t.Logf("a-out: %v\n", c.Data["num"])
	})
	m.Add(func(c *Context) {
		c.Data["num"] = c.Data["num"].(int) + 1
		t.Logf("b-in: %v\n", c.Data["num"])
		c.Next()
		t.Logf("b-out: %v\n", c.Data["num"])
	})
	m.Add(func(c *Context) {
		c.Data["num"] = c.Data["num"].(int) + 1
		t.Logf("c-in: %v\n", c.Data["num"])
		c.Next()
		t.Logf("c-out: %v\n", c.Data["num"])
	})
	m.Add(func(c *Context) {
		c.Data["num"] = c.Data["num"].(int) + 1
		t.Logf("d-in: %v\n", c.Data["num"])
		c.Next()
		t.Logf("d-out: %v\n", c.Data["num"])
	})

	t.Log("\n第一次flow")
	m.Flow(nil)

	t.Log("\n第二次flow")
	m.Flow(nil)

	t.Log("\n第三次flow")
	m.Flow(nil)
}

func TestCustomContext(t *testing.T) {
	m := NewApp()
	m.Use(func(c *CustomContext) {
		c.Data = 100
		t.Logf("a-in: %v\n", c.Data)
		c.Next()
		t.Logf("a-out: %v\n", c.Data)
	})
	m.Use(func(c *CustomContext) {
		c.Data = c.Data + 1
		t.Logf("b-in: %v\n", c.Data)
		c.Next()
		t.Logf("b-out: %v\n", c.Data)
	})
	m.Use(func(c *CustomContext) {
		c.Data = c.Data + 1
		t.Logf("c-in: %v\n", c.Data)
		c.Next()
		t.Logf("c-out: %v\n", c.Data)
	})
	m.Use(func(c *CustomContext) {
		c.Data = c.Data + 1
		t.Logf("d-in: %v\n", c.Data)
		c.Next()
		t.Logf("d-out: %v\n", c.Data)
	})

	t.Log("\n第一次flow")
	m.HanldeFlow()

	t.Log("\n第二次flow")
	m.HanldeFlow()

	t.Log("\n第三次flow")
	m.HanldeFlow()
}

type CustomContext struct {
	context *Context
	Next    func()
	Data    int8
}
type App struct {
	mw   *Application
	pool sync.Pool
	c    *CustomContext
}

func (a *App) Use(fn func(ctx *CustomContext)) {
	a.mw.Add(func(ctx *Context) {
		fn(a.c)
	})
}
func (a *App) HanldeFlow() {
	a.mw.Flow(func(ctx *Context) {
		c := a.pool.Get().(*CustomContext)
		c.context = ctx
		c.Next = ctx.Next
		c.Data = -1
		a.c = c
	})
	a.pool.Put(a.c)
}
func NewApp() *App {
	a := &App{
		mw: New(),
	}
	a.pool.New = func() interface{} {
		c := &CustomContext{Data: -1}
		return c
	}
	return a
}
