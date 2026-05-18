package handlers

import (
	"log"
	"net/http"
	"servidorHTTP/app/utils"
)

func RegistrarHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	nomeCompleto := request.FormValue("nomeCompleto")
	cpf := request.FormValue("cpf")
	dataNascimento := request.FormValue("dataNascimento")
	telefone := request.FormValue("telefone")
	email := request.FormValue("email")
	dataInternacao := request.FormValue("dataInternacao")
	horaInternacao := request.FormValue("horaInternacao")
	motivoInternacao := request.FormValue("motivoInternacao")

	log.Printf("Tentando registrar: %s, %s, %s", nomeCompleto, email, cpf)

	err := utils.InsertPaciente(nomeCompleto, cpf, dataNascimento, telefone, email, dataInternacao, horaInternacao, motivoInternacao)
	if err != nil {
		log.Printf("ERRO ao registrar paciente: %v", err)
		http.Error(response, "Erro ao registrar paciente", http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, "/menu.html", http.StatusSeeOther)
}
