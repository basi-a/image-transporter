package main

import (
	"fmt"
	"image-transporter/cmd"
	"image-transporter/global"
)

func main()  {
	defer global.Client.Close()
	list, err := cmd.List()
	if err != nil{
		panic(err)
	}
	fmt.Println(list)
}

func init()  {
	global.Init()
}