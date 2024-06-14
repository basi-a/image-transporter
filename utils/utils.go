package utils

import (
	"errors"
	"os"
	"regexp"
)
func DirExist(dir string) bool {
	_, err := os.Stat(dir)
	return errors.Is(err, os.ErrNotExist) 
}
func MkdirAll(dir string) error {
	return os.MkdirAll(dir, os.ModeDir)
}
func ReplaceUnsafeCharacters(str string) string {
	re := regexp.MustCompile(`[\\/]+`)
	safeStr := re.ReplaceAllString(str, "-")
	return safeStr
}