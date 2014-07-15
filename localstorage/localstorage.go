package localstorage

import (
	"os"
	"bytes"
	"io/ioutil"
	"path/filepath"
	"github.com/ShaneKilkelly/jetcan/config"
)

const DEFAULT_DIRECTORY	string = ".jetcan.d"
const AUTH_TOKEN_FILE = "token"

type LocalStorage struct {
	RootDir	string
}

func New(cfg *config.Config) (*LocalStorage, error) {
	l := &LocalStorage{cfg.StorageDir}
	err := l.Initialize()
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Initialize local storage
func (l *LocalStorage) Initialize() error {
	var err				error

	err = createStorageDir(l.RootDir)
	if err != nil {
		return err
	}

	return nil
}

func (l *LocalStorage) GetAuthToken() (string, error) {
	var (
		tokenFilePath	string
		content			[]byte
		token			string
		err				error
	)

	tokenFilePath, err = filepath.Abs(
		filepath.Join(l.RootDir, AUTH_TOKEN_FILE))
	if err != nil {
		return "", err
	}
	content, err = ioutil.ReadFile(tokenFilePath)
	if err != nil {
		return "", err
	}

	token = string(bytes.Trim(content, "\x00"))

	return token, nil
}

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
func createStorageDir(storageDir string) error {
	var (
		parentDir	string
		fullPath	string
		permissions	os.FileMode
		err			error
	)

	if storageDir != "" {
		storageDir, err = filepath.Abs(storageDir)
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

