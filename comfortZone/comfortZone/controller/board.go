package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Weltloose/comfortZone/model"

	"github.com/Weltloose/comfortZone/lib"
)

func Comment(w http.ResponseWriter, r *http.Request) {
	if !checkAuthed(r) {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	commentContent := struct {
		TextareaContents string `from:"PostForm" name:"textareaContents"`
	}{}
	err := lib.Unpack(r, &commentContent)
	if err != nil {
		fmt.Println("req unpack ", err)
		fmt.Fprintf(w, "false")
		return
	}
	if commentContent.TextareaContents == "" {
		fmt.Fprintf(w, "false")
		return
	}
	username := model.GetUsernameWithAuth(r.FormValue("czcookie"))
	if username != "" {
		if model.AddComment(username, commentContent.TextareaContents) {
			fmt.Fprintln(w, "true")
			return
		}
	}
	fmt.Fprintln(w, "false")

}

func GetComment(w http.ResponseWriter, r *http.Request) {
	if !checkAuthed(r) {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	lenStr := r.FormValue("saved")
	len, _ := strconv.Atoi(lenStr)
	fmt.Println("len", len)
	comments := model.FetchComment(len)
	data, _ := json.Marshal(comments)
	fmt.Fprintln(w, string(data))
}
