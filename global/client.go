package global

import (
	"context"
	"image-transporter/utils"

	"github.com/docker/docker/client"
)

var Client *client.Client
var Ctx context.Context
func InitDockerClient()  {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		utils.FatalfErr("init docker client fail.", err)
	}
	Client = client
	Ctx = context.Background()
}