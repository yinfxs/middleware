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

## Quick Start

```sh
# assume the following codes in example.go file
$ cat main.go
```

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
# run main.go
$ go run main.go
2018/06/25 18:23:47 handler a-in: 100
2018/06/25 18:23:47 handler b-in: 101
2018/06/25 18:23:47 handler c-in: 102
2018/06/25 18:23:47 handler c-out: 102
2018/06/25 18:23:47 handler b-out: 102
2018/06/25 18:23:47 handler a-out: 102
```

## Examples

### Define a handler

```go
func(c *middleware.Context) {
  c.Next()
}
```

### Extending Context

If you want to define a custom `Context`, fork it!

## Contributing

If you'd like to help out with the project. You can put up a Pull Request.
