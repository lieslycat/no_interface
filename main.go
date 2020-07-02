package no_interface

import (
		"fmt"
)

type No_Interface struct {
		Name string
}

func (c *No_Interface) Interface_func() {
  fmt.Printf("name: %v", c.name)
}
