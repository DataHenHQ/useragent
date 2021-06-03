# useragent
DataHen useragent is a go package and tool that generates a random combination of millions of user-agents strings. Currently used in production at DataHen to crawl/scrape through billions of pages.

# Usage

There are two ways to use this. Either as a golang package, or as a standalone binary.

## As Golang package

To use this as a golang package simply import it

```go
import github.com/DataHenHQ/useragent
```

Example code:

```go
package main

import (
	"fmt"

	"github.com/DataHenHQ/useragent"
)

func main() {
	fmt.Println("Trying out random user agents:")
	for i := 0; i < 20; i++ {
		ua, err := useragent.Desktop()
		if err != nil {
			panic(err)
		}
		fmt.Println(ua)
	}

}
```
You can see [more examples here](/examples)

## As standalone binary
To run as a standalone binary, [download](https://github.com/DataHenHQ/useragent/releases) the appropriate binary for your system, and install it.

You can then use it to generate a user agent string like so:

```shell
$ useragent generate desktop                                                    

# will output:
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36 Edg/85.0.564.44
```

If you would like to generate multiple combination of user-agents, you can do the following:

```shell
$ useragent generate desktop -n 2                                                    

# will output:
Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36 Edg/85.0.564.44
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147 Safari/537.36
```
