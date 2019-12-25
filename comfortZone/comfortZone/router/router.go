package router

import (
	"net/http"

	"github.com/Weltloose/comfortZone/controller"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// GetServer return web server
func GetServer() *negroni.Negroni {
	mux := mux.NewRouter()
	mux.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	mux.HandleFunc("/", controller.CheckSignedin).Methods("GET")
	mux.HandleFunc("/signIn", controller.SignIn).Methods("GET")
	mux.HandleFunc("/index", controller.Index).Methods("GET")
	mux.HandleFunc("/gallery", controller.Gallery).Methods("GET")
	mux.HandleFunc("/password", controller.Password).Methods("GET")
	mux.HandleFunc("/api/signInCheck", controller.SignInCheck).Methods("POST")
	mux.HandleFunc("/api/register", controller.Register).Methods("POST")
	mux.HandleFunc("/api/comment", controller.Comment).Methods("POST")
	mux.HandleFunc("/api/getComment", controller.GetComment).Methods("GET")
	mux.HandleFunc("/api/uploadPublicPhotoes", controller.UploadPublicPhotoes).Methods("POST")
	mux.HandleFunc("/api/uploadPublicVoice", controller.UploadPublicVoice).Methods("POST")
	mux.HandleFunc("/api/getPublicPhotoesInfo", controller.GetPublicPhotoesInfo).Methods("GET")
	mux.HandleFunc("/api/changePassword", controller.ChangePassword).Methods("PATCH")

	// mux.HandleFunc("/register", controller.Register)
	n := negroni.Classic()
	n.UseHandler(mux)
	return n
}
