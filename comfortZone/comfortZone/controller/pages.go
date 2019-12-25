package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("static/login.html")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(data))
	fmt.Fprintln(w, string(data))
}

func Gallery(w http.ResponseWriter, r *http.Request) {
	if !checkAuthed(r) {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	data, _ := ioutil.ReadFile("static/gallery.html")
	fmt.Fprintf(w, string(data))
}

func Index(w http.ResponseWriter, r *http.Request) {
	if !checkAuthed(r) {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	data, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, string(data))
}

func Password(w http.ResponseWriter, r *http.Request) {
	// if !checkAuthed(r) {
	// 	http.Redirect(w, r, "/signIn", http.StatusFound)
	// 	return
	// }
	data, _ := ioutil.ReadFile("static/passwordChange.html")
	fmt.Fprintf(w, string(data))
}
