package middleware

import (
	"sync"
)

// M is a shortcup for map[string]interface{}
type M map[string]interface{}

// Context 中间件上下文
type Context struct {
	Data  M
	index int8
	Next  func()
}

// Middleware 中间件类型
type Middleware struct {
	pool sync.Pool
	arr  []func(ctx *Context)
}

// Add 新增中间件
func (m *Middleware) Add(fn func(ctx *Context)) {
	m.arr = append(m.arr, fn)
}

// createContext 获取上下文对象
func (m *Middleware) createContext(data M) *Context {
	c := m.pool.Get().(*Context)
	if data == nil {
		data = M{}
	}
	c.Data = data
	c.index = -1
	return c
}

// Flow 流转中间件
func (m *Middleware) Flow(data M) {
	ctx := m.createContext(data)
	ctx.Next()
	m.pool.Put(ctx)
}

// New 创建中间件集合
func New() *Middleware {
	m := &Middleware{}
	m.pool.New = func() interface{} {
		ctx := &Context{
			index: -1,
			Data:  M{},
		}
		ctx.Next = func() {
			ctx.index++
			if ctx.index < int8(len(m.arr)) {
				fn := m.arr[ctx.index]
				if fn != nil {
					fn(ctx)
				}
			}
		}
		return ctx
	}
	return m
}
