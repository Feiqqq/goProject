package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Привет. Это приложение для организации эксклюзивных вечеринок. Перенаправляю на регистрацию</h1> "+
		"<meta http-equiv=\"Refresh\" content=\"5; URL=http://localhost:8080/user/registration\">")
}
