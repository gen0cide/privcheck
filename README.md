# Privcheck

A **very** small cross platform package that easily allows you to determine if you are executing with administrative or root privileges.

This library has zero external dependencies outside of the standard library and only has one function: `IsAdmin()` which returns a `bool`.

On nix systems, this is as simple as getting the effective UID. On Windows, an administrative handle is opened into the LSA's SCManager with requested administrative rights. This library then checks to see if Windows granted this request or not.

Very easy to use:

```golang
package main

import (
	"fmt"

	"github.com/gen0cide/privcheck"
)

func main() {
	if privcheck.IsAdmin() {
		fmt.Println("WE ARE ADMIN")
	} else {
		fmt.Println("WE ARE *NOT* ADMIN")
	}
}
```

This example can be found in `examples/main.go`.

Results:

```
$ ./example
WE ARE *NOT* ADMIN
$ sudo ./example
Password:
WE ARE ADMIN
$ 
```

Cheers!