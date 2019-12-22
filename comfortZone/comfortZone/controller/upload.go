package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/Weltloose/comfortZone/model"
)

func UploadPublicPhotoes(w http.ResponseWriter, r *http.Request) {
	respSuccess, _ := json.Marshal(struct {
		Resp string `json:"resp"`
	}{
		"success",
	})
	respFailed, _ := json.Marshal(struct {
		Resp string `json:"resp"`
	}{
		"failed",
	})
	if !checkAuthed(r) {
		fmt.Fprintln(w, respFailed)
		return
	}
	r.ParseMultipartForm(1024)
	fmt.Println(r.MultipartForm.File, len(r.MultipartForm.File))
	var tmpsSrc []string
	var realNames []string
	var destSrc []string
	for i := 0; i < len(r.MultipartForm.File); i++ {
		key := strconv.Itoa(i)
		fileHeader := r.MultipartForm.File[key][0]
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Fprintln(w, respFailed)
			return
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintln(w, respFailed)
			return
		}
		filePath := getTmpPath(fileHeader.Filename)
		realName := fileHeader.Filename
		toPath := getPublicPhotoesPath(filePath[strings.LastIndex(filePath, "/")+1:])
		err = ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			fmt.Fprintln(w, respFailed)
			return
		}
		tmpsSrc = append(tmpsSrc, filePath)
		realNames = append(realNames, realName)
		destSrc = append(destSrc, toPath)
	}
	for i := 0; i < len(tmpsSrc); i++ {
		model.TransferSource(tmpsSrc[i], destSrc[i], realNames[i])
	}

	fmt.Fprintln(w, string(respSuccess))
}

func UploadPublicVoice(w http.ResponseWriter, r *http.Request) {
	respSuccess, _ := json.Marshal(struct {
		Resp string `json:"resp"`
	}{
		"success",
	})
	respFailed, _ := json.Marshal(struct {
		Resp string `json:"resp"`
	}{
		"failed",
	})
	if !checkAuthed(r) {
		fmt.Fprintln(w, respFailed)
		return
	}
	r.ParseMultipartForm(1024)
	fmt.Println(r.MultipartForm.File, len(r.MultipartForm.File))
	var tmpsSrc []string
	var realNames []string
	var destSrc []string
	for i := 0; i < len(r.MultipartForm.File); i++ {
		key := strconv.Itoa(i)
		fileHeader := r.MultipartForm.File[key][0]
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Fprintln(w, respFailed)
			return
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintln(w, respFailed)
			return
		}
		filePath := getTmpPath(fileHeader.Filename)
		realName := fileHeader.Filename
		toPath := getPublicVoicePath(filePath[strings.LastIndex(filePath, "/")+1:])
		err = ioutil.WriteFile(filePath, data, 0644)
		if err != nil {
			fmt.Fprintln(w, respFailed)
			return
		}
		tmpsSrc = append(tmpsSrc, filePath)
		realNames = append(realNames, realName)
		destSrc = append(destSrc, toPath)
	}
	for i := 0; i < len(tmpsSrc); i++ {
		model.TransferSource(tmpsSrc[i], destSrc[i], realNames[i])
	}

	fmt.Fprintln(w, string(respSuccess))
}

func getTmpPath(fileName string) string {
	i := model.AddSource()
	file := strconv.Itoa(i)
	return "static/tmp/" + file
}

func getPublicPhotoesPath(fileName string) string {
	return "static/PublicPhotoes/" + fileName
}

func getPublicVoicePath(fileName string) string {
	return "static/PublicVoice/" + fileName
}
