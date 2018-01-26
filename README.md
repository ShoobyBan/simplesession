# simplesession

Really Simple Session management in Go, bob into file

```
func (SimpleSession) Load
func (s SimpleSession) Load(storageFilePath string) (SimpleSession, error)
Load loads decodes session from Gob file, saving will update the file

func (SimpleSession) Save
func (s SimpleSession) Save() error
Save encodes and saves session via Gob to file

func (SimpleSession) Set
func (s SimpleSession) Set(key string, val interface{})
Set sets a value into the store and saves session
```

Usage:
```go
// Simple counter program using session storage
package main
import (
    "github.com/shoobyban/simplesession"
    "fmt"
)
func main () {
    sess, err := simplesession.Load("sessionstore.data")
    if err == nil {
        z:= (*sess)["test"]
        if z != nil {
            fmt.Println(z.(int))
            sess.Set("test",z.(int)+1)
        } else {
            fmt.Println("Setting to 1, first run")
            sess.Set("test",1)
        }
    } else {
        fmt.Println(err)
    }
}
```

