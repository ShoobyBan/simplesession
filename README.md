# simplesession

Really Simple Session management in Go, bob into file

```
SimpleSession type, for session save/load
var StorageFilePath string
type SimpleSession map[string]interface{}
func Load(storageFilePath string) (*SimpleSession, error)
    Load loads decodes session from Gob file, saving will update the file

func (s *SimpleSession) Save() error
    Save encodes and saves session via Gob to file

func (s *SimpleSession) Set(key string, val interface{})
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

