# eval
A minimalistic math parser for Go. It implements the [shunting-yard-algorithm](https://brilliant.org/wiki/shunting-yard-algorithm/)
and allows to parse math from strings.

## Work in progress
The current implementation requires spaces between each math token and does not support trigionometric functions yet.

## Example
The library can be used as illustrated below:

```go
package main

import (
	"fmt"

	"github.com/jacalz/eval"
)

func main() {
	input := "( 6 - 2 * ( 6 / 3 ) ) ^ 3"
	
	result, err := eval.Evaluate(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
```

A more elaborate example can be found in the `example` folder.
