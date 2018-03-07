package helper

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// Exists returns whether the given file or directory exists or not.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// GitRootDir returns the Git root directory of the current project.
func GitRootDir() (string, error) {
	currDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for path.Clean(currDir) != "/" {
		isGitRootDir, fileStatErr := Exists(path.Join(currDir, ".git"))

		if fileStatErr != nil {
			return "", fileStatErr
		}

		if isGitRootDir {
			return currDir, nil
		}

		currDir = filepath.Dir(currDir)
	}

	return "", fmt.Errorf("No Git root accessible from the current working directory")
}
