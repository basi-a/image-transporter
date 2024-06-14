package cmd

import (
	"encoding/json"
	"image-transporter/global"

	"github.com/docker/docker/api/types/registry"
)
func Search(term string) (string, error)  {
	result, err := global.Client.ImageSearch(global.Ctx, term, registry.SearchOptions{})
	if err != nil {
		return "", err
	}else {
		jsondata, _ := json.Marshal(result)
		return string(jsondata), nil
	}
}