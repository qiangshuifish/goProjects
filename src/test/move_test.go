package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test_move_pic(t *testing.T) {
	// 指定要移动到的文件夹
	destDir := "/Users/wpy/Desktop/bi"

	// 获取当前目录下的所有子目录
	dirs, err := ioutil.ReadDir("/Users/wpy/Desktop/target")
	if err != nil {
		panic(err)
	}

	// 遍历每个子目录
	for _, dir := range dirs {
		if dir.IsDir() {
			fmt.Printf("Processing directory: %s\n", dir.Name())

			// 获取该子目录下的所有图片文件
			files, err := filepath.Glob(filepath.Join(dir.Name(), "*.jpg"))
			if err != nil {
				panic(err)
			}
			files = append(files, filepath.Join(dir.Name(), "*.png"))
			files = append(files, filepath.Join(dir.Name(), "*.gif"))
			files = append(files, filepath.Join(dir.Name(), "*.jpeg"))

			// 移动每个图片文件到指定文件夹中
			for _, file := range files {
				fmt.Printf("Moving file: %s\n", file)
				err = os.Rename(file, filepath.Join(destDir, filepath.Base(file)))
				if err != nil {
					panic(err)
				}
			}
		}
	}

	fmt.Println("Done.")
}
