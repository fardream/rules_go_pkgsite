package main

import (
	"fmt"

	"github.com/fardream/rules_go_pkgsite"
)

func main() {
	f := rules_go_pkgsite.Foo{
		ADouble: new(float64),
		AString: "abc",
	}

	fmt.Println(f.SomeOtherMethod())
}
