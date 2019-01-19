# gsema
a simple goroutine limit pool.

## Installation

```
$ go get github.com/EDDYCJY/gsema
```

## Usage

```
package main

import (
	"fmt"
	"time"

	"github.com/EDDYCJY/gsema"
)

var sema = gsema.NewSemaphore(3)

func main() {
	userCount := 10
	for i := 0; i < userCount; i++ {
		go Read(i)
	}

	sema.Wait()
}

func Read(i int) {
	defer sema.Done()
	sema.Add(1)

	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
}
```

## Output

```
...
go func: 0, time: 1547880322
go func: 2, time: 1547880322
go func: 1, time: 1547880322
go func: 3, time: 1547880323
go func: 4, time: 1547880323
go func: 5, time: 1547880323
go func: 7, time: 1547880324
go func: 8, time: 1547880324
go func: 6, time: 1547880324
go func: 9, time: 1547880325
```