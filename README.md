# Go Application Build Information

This package is heavily work in progress. Right now, you have zero options to
customise the things, long term goals is to provide a much better experience
with configurable parts.


## Example

```go
package main

import (
	"fmt"

	"github.com/go-asset/build"
)

const AppName = "MyFancyApp"

func main() {
	version, _ := build.ReadVersion(AppName)
	fmt.Println(version)
}
```

```
❯ go run  -ldflags="-buildid=v0.0.1" .
MyFancyApp v0.0.1 *dirty* (2022-04-04T11:05:34Z)
```

