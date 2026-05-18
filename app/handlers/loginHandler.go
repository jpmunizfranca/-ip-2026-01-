package handlers

import (
	"net/http"
	"text/template"
)

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	username := request.FormValue("username")
	password := request.FormValue("password")

	if username == "" || password == "" {
		http.Error(response, "Usuário ou senha vazios", http.StatusBadRequest)
		return
	}

	tmpl, err := template.ParseFiles("static/menu.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(response, nil)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}
}
