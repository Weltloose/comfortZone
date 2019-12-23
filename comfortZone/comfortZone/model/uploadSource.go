package model

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/nickalie/go-webpbin"
)

func AddSource() int {
	tmp := GetCurSources()
	addCurSources()
	return tmp
}

func TransferSource(sourceFile, destFile, realFile string) {
	file, _ := os.Open(sourceFile)
	defer file.Close()
	img, imgType, _ := image.Decode(file)
	if imgType != "" {
		f, _ := os.Create(destFile)
		defer f.Close()
		if err := webpbin.Encode(f, img); err != nil {
			return
		}
		newSource(destFile, realFile)
	} else {
		data, _ := ioutil.ReadFile(sourceFile)
		if img, err := webpbin.Decode(bytes.NewReader(data)); err == nil {
			f, _ := os.Create(destFile)
			defer f.Close()
			if err := webpbin.Encode(f, img); err != nil {
				f.Close()
				return
			}
			newSource(destFile, realFile)
		} else {
			os.Remove(destFile)
			fmt.Println("sourceFile error ", err)
		}
	}

}
