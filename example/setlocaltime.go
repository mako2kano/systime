// Copyright 2018 Makoto Tokano.
// Use of this source code is governed by a zlib-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/mako2kano/systime"
)

func main() {
	t := time.Now()
	err := systime.SetLocalTime(&t)

	fmt.Println(err)
}
