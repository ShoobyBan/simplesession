package simplesession

import (
	"encoding/gob"
	"os"
)

// SimpleSession type, for session save/load
type SimpleSession struct {
	storageFilePath string
	m               map[string]interface{}
}

// Save encodes and saves session via Gob to file
func (s SimpleSession) Save() error {
	file, err := os.Create(s.storageFilePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(s)
	}
	file.Close()
	return err
}

// Load loads decodes session from Gob file, saving will update the file
func (s SimpleSession) Load(storageFilePath string) (SimpleSession, error) {
	file, err := os.Open(storageFilePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(s)
	}
	s.storageFilePath = storageFilePath
	file.Close()
	return s, err
}

// Set sets a value into the store and saves session
func (s SimpleSession) Set(key string, val interface{}) {
	s.m[key] = val
	s.Save()
}
