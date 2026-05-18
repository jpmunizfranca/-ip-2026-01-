package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
	"text/template"
)

func EditarPageHandler(response http.ResponseWriter, request *http.Request) {
	idStr := request.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID inválido", http.StatusBadRequest)
		return
	}

	paciente, err := utils.GetPacienteByID(id)
	if err != nil {
		http.Error(response, "Paciente não encontrado", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/forms/editar.html")
	if err != nil {
		http.Error(response, "Erro ao carregar o template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(response, paciente)
	if err != nil {
		http.Error(response, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}
}

func EditarHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	idStr := request.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(response, "ID inválido", http.StatusBadRequest)
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

	err = utils.UpdatePaciente(id, nomeCompleto, cpf, dataNascimento, telefone, email, dataInternacao, horaInternacao, motivoInternacao)
	if err != nil {
		http.Error(response, "Erro ao atualizar paciente", http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, "/listar", http.StatusSeeOther)
}
