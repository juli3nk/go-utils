package filedir

import (
	"io/ioutil"
	"os"
	"path"
)

func DirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateDirIfNotExist(path string, parent bool, perm os.FileMode) error {
	if !DirExists(path) {
		if parent {
			if err := os.MkdirAll(path, perm); err != nil {
				return err
			}
		} else {
			if err := os.Mkdir(path, perm); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyDir(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return err
	}

	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, obj := range files {
		fsource := path.Join(src, obj.Name())
		fdest := path.Join(dst, obj.Name())

		if obj.IsDir() {
			// Create sub-directories, recursively.
			if err := CopyDir(fsource, fdest); err != nil {
				return err
			}
		} else {
			if err := CopyFile(fsource, fdest); err != nil {
				return err
			}
		}
	}

	return nil
}
