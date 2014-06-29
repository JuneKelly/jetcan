package localstorage

import (
	"os"
	"os/user"
	"path/filepath"
)

const DIRECTORY	string = ".jetcan.d"

// helper, get the FileMode for a path
func getMode(path string) (os.FileMode, error) {
	var fi	os.FileInfo
	var err	error

	fi, err = os.Stat(path)
	if err != nil {
		return 0, err
	}

	return fi.Mode(), nil
}

// Ensure that the storage directory exists under the
// current users home directory
func createStorageDir() error {
	var currentUser	*user.User
	var path		string
	var permissions	os.FileMode
	var err			error

	currentUser, err = user.Current()
	if err != nil {
		return err
	}

	path = filepath.Join(currentUser.HomeDir, DIRECTORY)
	permissions, err = getMode(currentUser.HomeDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path, permissions)
	if err != nil {
		return err
	}

	return nil
}

// Initialize local storage
func Initialize() error {
	var err				error

	err = createStorageDir()
	if err != nil {
		return err
	}

	return nil
}
