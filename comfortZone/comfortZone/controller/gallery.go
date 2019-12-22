package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Weltloose/comfortZone/model"
)

func Gallery(w http.ResponseWriter, r *http.Request) {
	if !checkAuthed(r) {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	data, _ := ioutil.ReadFile("static/gallery.html")
	fmt.Fprintf(w, string(data))
}

func GetPublicPhotoesInfo(w http.ResponseWriter, r *http.Request) {
	if !checkAuthed(r) {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	opt, _ := json.Marshal(model.GetPublicPhotoesInfo())
	fmt.Fprintf(w, string(opt))
}
