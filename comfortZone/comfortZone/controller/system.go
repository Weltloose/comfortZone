package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Weltloose/comfortZone/model"
)

func CheckSignedin(w http.ResponseWriter, r *http.Request) {
	if checkAuthed(r) {
		fmt.Fprintf(w, "success")
		return
	}
	http.Redirect(w, r, "/signIn", http.StatusFound)
}

func SignInCheck(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
	}{}
	data, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(data, &req)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	}
	fmt.Println("username: ", req.Username, req.Passwd)
	signedIn, val := model.CheckValid(req.Username, req.Passwd)
	if !signedIn {
		fmt.Println("Username and passwd: ", req.Username, req.Passwd)
		fmt.Fprintf(w, "Invalid")
		return
	}
	fmt.Fprintf(w, "success%s", val)
}

func Register(w http.ResponseWriter, r *http.Request) {
	req := struct {
		Username string `json:"username"`
		Passwd   string `json:"passwd"`
	}{}
	data, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(data, &req)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	}
	if model.CheckExist(req.Username) {
		fmt.Fprintln(w, "username already exist")
		return
	}
	if !model.RegisterUser(req.Username, req.Passwd) {
		fmt.Fprintln(w, "register error")
		return
	}
	fmt.Fprintln(w, "register success, now try sign in!")
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	if !checkAuthed(r) {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	fmt.Println(r.Form)
	fmt.Println(r.PostForm)
	password := r.PostFormValue("passwd")
	username := model.GetUsernameWithAuth(r.FormValue("czcookie"))
	ok := model.ChangePassword(username, password)
	if ok {
		fmt.Fprintln(w, "password change success!")
	}
}

func checkAuthed(r *http.Request) bool {
	c := r.FormValue("czcookie")
	fmt.Println("authed: ", c)
	logined, err := model.CheckLogined(c)
	if err == nil && logined {
		// http.Redirect(w, r, "/index", http.StatusFound)
		return true
	}
	fmt.Println("dude you cann't")

	fmt.Println(err)
	return false
}
