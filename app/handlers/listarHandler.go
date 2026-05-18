package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"text/template"
)

func ListarHandler(response http.ResponseWriter, request *http.Request) {
	pacientes, err := utils.GetAllPacientes()
	if err != nil {
		http.Error(response, "Erro ao buscar pacientes", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/listar.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(response, pacientes)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}
}
