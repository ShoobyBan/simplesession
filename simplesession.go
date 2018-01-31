package simplesession

import (
	"encoding/gob"
	"os"
	"sync"
)

// SimpleSession type, for session save/load
var StorageFilePath string

type SimpleSession map[string]interface{}

var lock sync.Mutex

// Save encodes and saves session via Gob to file
func (s *SimpleSession) Save() error {
	file, err := os.Create(StorageFilePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(s)
	}
	file.Close()
	return err
}

// Load loads decodes session from Gob file, saving will update the file
func Load(storageFilePath string) (*SimpleSession, error) {
	file, err := os.Open(storageFilePath)
	s := new(SimpleSession)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(s)
	} else {
		n := make(SimpleSession, 1)
		s = &n
		err = nil
	}
	StorageFilePath = storageFilePath
	file.Close()
	return s, err
}

// Set sets a value into the store and saves session
func (s *SimpleSession) Set(key string, val interface{}) {
	lock.Lock()
	defer lock.Unlock()
	(*s)[key] = val
	s.Save()
}
