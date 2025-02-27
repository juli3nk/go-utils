package filedir

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func FileExists(f string) bool {
	if _, err := os.Lstat(f); err != nil {
		return false
	}

	return true
}

func IsSymlink(f string) (bool, string, error) {
	var t bool
	var link string

	fi, err := os.Lstat(f)
	if err != nil {
		return t, link, err
	}

	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		t = true
		link, err = os.Readlink(f)
		if err != nil {
			return t, link, err
		}
	}

	return t, link, nil
}

func CopyFile(src, dst string) error {
	sfi, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sfi.Mode().IsRegular() {
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	dfi, err := os.Stat(dst)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}

		if os.SameFile(sfi, dfi) {
			return nil
		}
	}

	if err = copyFileContents(src, dst); err != nil {
		return err
	}

	return nil
}

func copyFileContents(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err = io.Copy(out, in); err != nil {
		return err
	}

	if err = out.Sync(); err != nil {
		return err
	}

	return nil
}

func FilterFileByExtension(files []string, ext string) []string {
	if ext == "" {
		return nil
	}
	ext = "." + ext

	var result []string

	for _, f := range files {
		if strings.HasSuffix(strings.ToLower(f), strings.ToLower(ext)) {
			result = append(result, f)
		}
	}

	return result
}
