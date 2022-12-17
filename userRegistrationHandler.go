package main

import (
	"html/template"
	"net/http"
)

func userRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	url := "html/registration.html" //Страница входа
	if len(r.Header["Cookie"]) != 0 && r.Header["Cookie"][0] == "auth=your_MD5_cookies" {
		url = "html/index.html" //Страница после успешной авторизации
	}
	t, _ := template.ParseFiles(url)
	t.Execute(w, "")
}
