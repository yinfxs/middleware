# Middleware

Go Middleware Framework

## Installation

```sh
go get github.com/yinfxs/middleware
```

## Import

```go
import "github.com/yinfxs/middleware"
```

## Define a middleware

```go
func(c *middleware.Context) {
  c.Next()
}
```

## Usage

```go
package main

import (
  "log"

  "github.com/yinfxs/middleware"
)

func main() {
  m := middleware.New()

  // handler a
  m.Add(func(c *middleware.Context) {
    c.Data["num"] = 100
    log.Printf("handler a-in: %v\n", c.Data["num"])
    c.Next()
    log.Printf("handler a-out: %v\n", c.Data["num"])
  })

  // handler b
  m.Add(func(c *middleware.Context) {
    c.Data["num"] = c.Data["num"].(int) + 1
    log.Printf("handler b-in: %v\n", c.Data["num"])
    c.Next()
    log.Printf("handler b-out: %v\n", c.Data["num"])
  })

  // handler c
  m.Add(func(c *middleware.Context) {
    c.Data["num"] = c.Data["num"].(int) + 1
    log.Printf("handler c-in: %v\n", c.Data["num"])
    c.Next()
    log.Printf("handler c-out: %v\n", c.Data["num"])
  })

  m.Flow(nil)
}
```

```sh
$ go run main.go
2018/06/25 18:23:47 handler a-in: 100
2018/06/25 18:23:47 handler b-in: 101
2018/06/25 18:23:47 handler c-in: 102
2018/06/25 18:23:47 handler c-out: 102
2018/06/25 18:23:47 handler b-out: 102
2018/06/25 18:23:47 handler a-out: 102
```

## Custom Context

```go
package main

import (
  "log"
  "sync"

  "github.com/yinfxs/middleware"
)

// CustomContext 自定义上下文对象
type CustomContext struct {
  context *middleware.Context
  Next    func()
  Data    int8
}

// App 应用
type App struct {
  mw   *middleware.Application
  pool sync.Pool
  c    *CustomContext
}

// Use 应用
func (a *App) Use(fn func(ctx *CustomContext)) {
  a.mw.Add(func(ctx *middleware.Context) {
    fn(a.c)
  })
}

// HanldeFlow 应用
func (a *App) HanldeFlow() {
  a.mw.Flow(func(ctx *middleware.Context) {
    c := a.pool.Get().(*CustomContext)
    c.context = ctx
    c.Next = ctx.Next
    c.Data = -1
    a.c = c
  })
  a.pool.Put(a.c)
}

// New 创建应用实例
func New() *App {
  a := &App{
    mw: middleware.New(),
  }
  a.pool.New = func() interface{} {
    c := &CustomContext{Data: -1}
    return c
  }
  return a
}

func main() {
  m := New()
  // handler a
  m.Use(func(c *CustomContext) {
    c.Data = 100
    log.Printf("handler a-in: %v\n", c.Data)
    c.Next()
    log.Printf("handler a-out: %v\n", c.Data)
  })
  // handler b
  m.Use(func(c *CustomContext) {
    c.Data = c.Data + 1
    log.Printf("handler b-in: %v\n", c.Data)
    c.Next()
    log.Printf("handler b-out: %v\n", c.Data)
  })
  // handler c
  m.Use(func(c *CustomContext) {
    c.Data = c.Data + 1
    log.Printf("handler c-in: %v\n", c.Data)
    c.Next()
    log.Printf("handler c-out: %v\n", c.Data)
  })
  m.HanldeFlow()
}
```

## Contributing

If you'd like to help out with the project. You can put up a Pull Request.
