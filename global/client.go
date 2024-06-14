package global

import (
	"context"

	"github.com/docker/docker/client"
)

var Client *client.Client
var Ctx context.Context
func InitDockerClient()  {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	Client = client
	Ctx = context.Background()
}