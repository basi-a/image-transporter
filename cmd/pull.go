package cmd

import (
	"fmt"
	"image-transporter/global"
	"io"
	"time"

	"github.com/docker/docker/api/types/image"
)

func Pull(refStr string) (string, time.Duration, error) {
	startTime := time.Now()
	readCloser, err := global.Client.ImagePull(global.Ctx, refStr, image.PullOptions{})
	if err != nil {
		return refStr, -1, err
	}
	defer readCloser.Close()
	// Process the pull output (optional)
    _, err = io.Copy(io.Discard, readCloser)
    if err != nil {
        return refStr, -1, fmt.Errorf("error reading pull output: %w", err)
    }
	pullTime := time.Since(startTime)
	return refStr, pullTime, nil
}