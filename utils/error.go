package utils

import "log"
func PrintfErr(str string, err error)  {
	log.Printf("%s : %s", str, err)
}
func FatalfErr(str string, err error)  {
	log.Fatalf("%s : %s", str, err)
}