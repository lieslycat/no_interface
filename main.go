package no_interface

import (
		"fmt"
)

type No_Interface struct {
		Name string
}

func (c *No_Interface) interface_func() {
  fmt.Printf("name: %v", c.name)
}
