package middleware

import (
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
