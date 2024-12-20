// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hi is a example pkg.
package hi

import (
	"fmt"

	"github.com/Hunted-Labs/gettext-go"
)

func SayHi() {
	fmt.Println(gettext.Gettext("pkg hi: Hello, world!"))
	fmt.Println(gettext.PGettext("code.google.com/p/gettext-go/examples/hi.SayHi", "pkg hi: Hello, world!"))
}
