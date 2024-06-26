package utils

import (
	"os"
	"regexp"
)
func DirExist(dir string) bool {
	if _, err := os.Stat(dir); err != nil {
		return false
	}
	return true 
}
func MkdirAll(dir string) error {
	return os.MkdirAll(dir, 0755)
}
func FileStat(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true 
}
func ReplaceUnsafeCharacters(str string) string {
	re := regexp.MustCompile(`[\\/]+`)
	safeStr := re.ReplaceAllString(str, "-")
	return safeStr
}
