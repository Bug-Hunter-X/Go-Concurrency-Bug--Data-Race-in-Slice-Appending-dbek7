```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var data []int
        var m sync.RWMutex

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m.Lock()
                        data = append(data, i)
                        m.Unlock()
                }(i)
        }

        wg.Wait()
        fmt.Println(len(data))
}
```