package model

import (
	"io/ioutil"
	"sort"
	"strconv"
)

type PublicPhotoInfo struct {
	IdName   string `json:"idName"`
	RealName string `json:"realName"`
	PostTime string `json:"postTime"`
	Liked    int    `json:"liked"`
}

type TotPublicPhotoesInfo struct {
	Data []PublicPhotoInfo `json:"data"`
}

func GetPublicPhotoesInfo() TotPublicPhotoesInfo {
	publicPhotoesFiles := getPublicPhotoesFiles()
	opt := TotPublicPhotoesInfo{}
	for _, file := range publicPhotoesFiles {
		opt.Data = append(opt.Data, readSource("static/PublicPhotoes/"+file))
	}
	return opt
}

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}
func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntSlice) Less(i, j int) bool {
	return s[i] > s[j]
}

func getPublicPhotoesFiles() []string {
	files, _ := ioutil.ReadDir("static/PublicPhotoes")
	opt := make([]string, len(files))
	tmp := make([]int, len(files))
	for i, file := range files {
		tint, _ := strconv.Atoi(file.Name())
		tmp[i] = tint
	}
	sort.Sort(IntSlice(tmp))
	for i, val := range tmp {
		opt[i] = strconv.Itoa(val)
	}
	return opt
}
