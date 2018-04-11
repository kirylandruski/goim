package main

import (
	"goim/auth"
	"os"
)

// extends User manager with methods allowing to read and write a file
type fileUserManager struct {
	auth.UserManager
	fileName string
}

func NewFileUserManager(fileName string) *fileUserManager {
	return &fileUserManager{UserManager: *auth.NewUserManager(), fileName: fileName}
}

func (m *fileUserManager) read() error {
	source, err := os.OpenFile(m.fileName, os.O_RDONLY|os.O_CREATE, 0600)
	defer source.Close()

	if err != nil {
		return err
	}
	err = m.Read(source)
	if err != nil {
		return err
	}

	return nil
}

func (m *fileUserManager) write() error {
	destination, err := os.OpenFile(m.fileName, os.O_WRONLY|os.O_CREATE, 0600)
	defer destination.Close()

	if err != nil {
		return err
	}
	destination.Truncate(0)
	err = m.Write(destination)
	if err != nil {
		return err
	}

	return nil
}
