package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/picspider/pkg/setting"
)

// DownloadPhoto download photo
func DownloadPhoto(dirname, title, band string) {
	imgDir := path.Join(setting.AppSetting.Path, dirname)
	file, err := os.Stat(imgDir)
	if err != nil || !file.IsDir() {
		dirErr := os.Mkdir(imgDir, os.ModePerm)
		if dirErr != nil {
			fmt.Printf("create dir %q faild\n", imgDir)
			os.Exit(1)
		}
	}

	log.Printf("Get %s", band)
	u, err := url.Parse(band)
	if err != nil {
		log.Println("parse url failed:", band, err)
		return
	}

	tmp := strings.TrimLeft(u.Path, "/")
	tmp = strings.ToLower(strings.Replace(tmp, "/", "-", -1))
	dotIndex := strings.LastIndex(tmp, ".")
	name := title + tmp[dotIndex:]
	filename := path.Join(imgDir, name)

	if checkExists(filename) {
		log.Printf("Exists %s", filename)
		return
	}

	response, err := http.Get(band)
	if err != nil {
		log.Println("get band failed:", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", band, err)
		return
	}

	image, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}

	defer image.Close()
	image.Write(data)
}

func checkExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
