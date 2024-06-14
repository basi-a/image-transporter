package main

import (
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
	result, err := cmd.Search("golang")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func init()  {
	global.Init()
}