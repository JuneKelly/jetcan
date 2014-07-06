package localstorage

import (
	"os"
	"path/filepath"
	"github.com/ShaneKilkelly/jetcan/config"
)

const DEFAULT_DIRECTORY	string = ".jetcan.d"

// helper, get the FileMode for a path
func getMode(path string) (mode os.FileMode, err error) {
	var fi	os.FileInfo

	fi, err = os.Stat(path)
	if err != nil {
		return 0, err
	}

	mode = fi.Mode()

	return
}

// Ensure that the storage directory exists under the
// current users home directory
func createStorageDir(cfg *config.Config) error {
	var (
		storageDir	string
		parentDir	string
		fullPath	string
		permissions	os.FileMode
		err			error
	)

	if cfg.StorageDir != "" {
		storageDir, err = filepath.Abs(cfg.StorageDir)
		if err != nil {
			return err
		}
	} else {
		storageDir = DEFAULT_DIRECTORY
	}

	fullPath = filepath.Join(storageDir)
	parentDir = filepath.Join(storageDir, "..")

	// this is a poor hack to just assign the mode of the parent dir
	// which will work presuming the user has correct
	// permissions on the parent dir.
	// TODO: figure out what permissions this dir should have
	permissions, err = getMode(parentDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(fullPath, permissions)
	if err != nil {
		return err
	}

	return nil
}

// Initialize local storage
func Initialize(cfg *config.Config) error {
	var err				error

	err = createStorageDir(cfg)
	if err != nil {
		return err
	}

	return nil
}
