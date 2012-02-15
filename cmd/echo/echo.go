package main

import (
	"github.com/nu7hatch/cgoecho2/pkg/echo"
	"flag"
)

func main() {
	flag.Parse()
	echo.Echo(flag.Args()...)
}