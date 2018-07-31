package main

import (
	"fmt"
	"time"

	"github.com/mako2kano/systime"
)

func main() {
	t := time.Now()
	err := systime.SetLocalTime(t)

	fmt.Println(err)
}
