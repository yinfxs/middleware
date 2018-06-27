package middleware

import (
	"sync"
)

// M is a shortcup for map[string]interface{}
type M map[string]interface{}

// Context 中间件上下文
type Context struct {
	Data     M
	Handlers []func(ctx *Context)
	index    int8
	Next     func()
}

// Application 中间件应用
type Application struct {
	pool sync.Pool
	arr  []func(ctx *Context)
}

// Add 新增中间件
func (m *Application) Add(fn func(ctx *Context)) {
	m.arr = append(m.arr, fn)
}

// createContext 获取上下文对象
func (m *Application) createContext() *Context {
	c := m.pool.Get().(*Context)
	c.Data = M{}
	c.index = -1
	return c
}

// Flow 流转中间件
func (m *Application) Flow(ctxReceiver func(ctx *Context)) {
	ctx := m.createContext()
	if ctxReceiver != nil {
		ctxReceiver(ctx)
	}
	ctx.Next()
	m.pool.Put(ctx)
}

// New 创建中间件应用
func New() *Application {
	m := &Application{}
	m.pool.New = func() interface{} {
		ctx := &Context{
			index:    -1,
			Data:     M{},
			Handlers: m.arr[:],
		}
		ctx.Next = func() {
			ctx.index++
			if ctx.index < int8(len(ctx.Handlers)) {
				fn := ctx.Handlers[ctx.index]
				if fn != nil {
					fn(ctx)
				}
			}
		}
		return ctx
	}
	return m
}
