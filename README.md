![release-beta](https://rawgit.com/suite911/assets/master/shields/release-beta-yellowgreen.svg)
![api-candidate](https://rawgit.com/suite911/assets/master/shields/api-candidate-green.svg)
![abi-alpha](https://rawgit.com/suite911/assets/master/shields/abi-alpha-orange.svg)
[![Build Status](https://travis-ci.org/suite911/env911.svg?branch=master)](https://travis-ci.org/suite911/env911)
[![CC0-1.0](https://rawgit.com/suite911/assets/master/shields/license-cc0--1.0-efbfff.svg)](https://raw.githubusercontent.com/suite911/cloud911/master/LICENSE.txt)

# env911
A [free](https://creativecommons.org/publicdomain/zero/1.0/) configuration library for Google Go (golang).

## API Stability
- app ![api-candidate](https://rawgit.com/suite911/assets/master/shields/api-candidate-green.svg)
- config ![api-candidate](https://rawgit.com/suite911/assets/master/shields/api-candidate-green.svg)
- safesave ![api-stable](https://rawgit.com/suite911/assets/master/shields/api-stable-brightgreen.svg)

Assume any packages not listed in the above are pre-alpha.

## License
This library is [truly free](https://creativecommons.org/publicdomain/zero/1.0/), not GNU-encumbered "free software".  Use it for any purpose, even commercial.  Attribution is appreciated but not required.

## Installation
```sh
go get -u github.com/suite911/env911/...
```

## Imports
```go
import "github.com/suite911/env911/app"
import "github.com/suite911/env911/config"
import "github.com/suite911/env911/safesave"
```

## Usage Examples
```go
package main

import (
	"fmt"

	"github.com/suite911/env911"
	"github.com/suite911/env911/app"
	"github.com/suite911/env911/config"
)

func init() {
	env911.InitAll("MYAPP_", nil, "MyCompany", "myapp")
}

func main() {
	verbose := false
	config.BoolVarP(&verbose, "verbose", "v", false, "Use verbose mode")
	config.LoadAndParse()
	if verbose {
		fmt.Printfn("Vendor:", app.Vendor()) // prints "MyCompany"
		fmt.Printfn("App:   ", app.Name()) // prints "myapp"
		fmt.Printfn("Path:  ", app.Path()) // prints "MyCompany/myapp" on POSIX systems
	}
}
```
