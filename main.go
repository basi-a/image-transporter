package main

import (
	// "fmt"
	"fmt"
	"image-transporter/cmd"
	"image-transporter/global"
)

func main()  {
	defer global.Client.Close()
	// list, err := cmd.List()
	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Println(list)

	// result, err := cmd.Search("golang")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)

	refStr, pullTime, err := cmd.Pull("golang:latest")
	if  err != nil {
		panic(err)
	}
	fmt.Println(refStr, pullTime)
}

func init()  {
	global.Init()
}