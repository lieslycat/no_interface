package no_interface

import (
		"fmt"
)

type No_Interface struct {
		name string
}

func (c *No_Interface) interface_func() {
  fmt.Printf("name: %v", c.name)
}
