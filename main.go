package main

import (
	"encoding/json"
	"fmt"
	"image-transporter/global"
	"image-transporter/initialize"
	"image-transporter/model"
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

	// refStr, pullTime, err := cmd.Pull("golang:latest")
	// if  err != nil {
	// 	panic(err)
	// }
	// fmt.Println(refStr, pullTime)
	
	// if err := cmd.Save([]string{"golang:latest", "nginx:latest", "nsqio/nsq:latest"}, "/home/yang/image-transporter"); err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(global.Config)
	
	result, err := model.UsersList(500)
	if err != nil {
		panic(err)
	}
	jsondata, _ := json.Marshal(result)
	fmt.Println(string(jsondata))
}

func init()  {
	initialize.Init()
}