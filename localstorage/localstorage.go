package localstorage

import (
	"os"
	"os/user"
	"path/filepath"
	"fmt"
)

const DIRECTORY	string = ".jetcan.d"

func userHomeDir() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	} else {
		return u.HomeDir, nil
	}
}

func exists() (bool, error) {
	var homeDir	string
	var path	string
	var err		error

	homeDir, err = userHomeDir()
	if err != nil {
		return false, err
	}

	path = filepath.Join(homeDir, DIRECTORY)

    _, err = os.Stat(path)
    if err == nil {
		return true, nil
	}
    if os.IsNotExist(err) {
		return false, nil
	}
    return false, err
}

func createStorageDir() error {
	var homeDir	string
	var path	string
	var err		error

	homeDir, err = userHomeDir()
	if err != nil {
		return err
	}
	path = filepath.Join(homeDir, DIRECTORY)

	fmt.Println("Creating local storage direcotory", path)

	err = os.Mkdir(path, os.ModeDir)
	if err != nil {
		return err
	}

	return nil
}

func Initialize() error {
	var storageExists	bool
	var err				error

	storageExists, err = exists()
	if err != nil {
		return err
	}

	if !storageExists {
		err = createStorageDir()
		if err != nil {
			return err
		}
	}
	return nil
}
