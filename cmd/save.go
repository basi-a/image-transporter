package cmd

import (
	"compress/gzip"
	"fmt"
	"image-transporter/global"
	"image-transporter/utils"
	"io"
	"os"
	"path/filepath"

)

func Save(imageIDs []string, outputDir string) error {
	for _, imageID := range imageIDs {
		readCloser, err := global.Client.ImageSave(global.Ctx, []string{imageID})
		if err != nil {
			return fmt.Errorf("failed to save image %s: %v", imageID, err)
		}
		defer readCloser.Close()
		if !utils.DirExist(outputDir){
			if err := utils.MkdirAll(outputDir); err != nil {
				return err
			}
		}
		
		outputFile := filepath.Join(outputDir, fmt.Sprintf("%s.tar.gz", utils.ReplaceUnsafeCharacters(imageID)))
		err = saveAndCompress(readCloser, outputFile)
		if err != nil {
			return fmt.Errorf("failed to compress image %s: %v", imageID, err)
		}
	}
	return nil
}

func saveAndCompress(reader io.Reader, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", outputFile, err)
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, reader)
	if err != nil {
		return fmt.Errorf("failed to write to gzip file %s: %v", outputFile, err)
	}

	return nil
}
