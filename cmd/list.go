package cmd

import (
	"encoding/json"
	"image-transporter/global"

	"github.com/docker/docker/api/types/image"
)
func List() (string, error) {
	summary, err := global.Client.ImageList(global.Ctx, image.ListOptions{
		All: true,
	})
	
	if err != nil {
		return "", err
	}else {
		jsondata, _ := json.Marshal(summary)
		return string(jsondata), nil
	}
}
