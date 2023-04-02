package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	sourceDir := "/Users/wpy/Downloads"

	targetDir := "/Users/wpy/Downloads/FC2PPV"
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		log.Fatalf("Target directory does not exist: %v", err)
	}

	regexString := ".*\\.mp4" // 正则表达式
	regex, err := regexp.Compile(regexString)
	if err != nil {
		log.Fatalf("Failed to compile regular expression: %v", err)
	}

	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Failed to access path %q: %v", path, err)
		}
		if !info.IsDir() && regex.MatchString(info.Name()) && info.Size() >= 1024*1024*100 {
			oldPath := path
			newPath := filepath.Join(targetDir, info.Name())
			if err := os.Rename(oldPath, newPath); err != nil {
				return fmt.Errorf("Failed to move file from %q to %q: %v", oldPath, newPath, err)
			}
			fmt.Printf("Moved file %s to %s\n", oldPath, newPath)
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error occurred while processing directory: %v", err)
	}
}
