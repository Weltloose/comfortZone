package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Weltloose/comfortZone/model"
)

func GetPublicPhotoesInfo(w http.ResponseWriter, r *http.Request) {
	if !checkAuthed(r) {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	opt, _ := json.Marshal(model.GetPublicPhotoesInfo())
	fmt.Fprintf(w, string(opt))
}
