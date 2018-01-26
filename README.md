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
```
  var sess limsession.LIMSession
	sess.Load("sessionstore.data")
  sess["test"] = "something"
  sess.Save()
```

