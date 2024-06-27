package utils

import (
	"os"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
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

func RandomStr(length int) string {
	const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+=-()"
	seededRand := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	str := make([]byte, length)
	for i := range str {
		charsetLength := len(charset)
		randomIndex := seededRand.Intn(charsetLength)
		randomCharByte := charset[randomIndex]
		str[i] = randomCharByte
	}
	return string(str)
}

func HashPassword(password string) (hashedPassword string, err error) {
	hashedstr, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword = string(hashedstr)
	return
}

func VerifyPassword(password, hashedPassword string) (ok bool, err error)  {
	bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return
}