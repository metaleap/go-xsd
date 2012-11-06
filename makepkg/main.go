package main

import (
	"fmt"

	xsd "github.com/metaleap/go-xsd"
)

func main () {
	var sd *xsd.Schema
	var err error
	if sd, err = xsd.LoadSchemaFile("C:\\gd\\src\\github.com\\metaleap\\_misc\\xsd\\coll14.xsd"); err != nil {
		panic(err)
	} else {
		fmt.Printf("%#v", sd)
	}
}
