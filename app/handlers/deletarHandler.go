package handlers

import (
	"net/http"
	"servidorHTTP/app/utils"
	"strconv"
)

func DeletarHandler(response http.ResponseWriter, request *http.Request) {
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

	err = utils.DeletePaciente(id)
	if err != nil {
		http.Error(response, "Erro ao deletar paciente", http.StatusInternalServerError)
		return
	}

	http.Redirect(response, request, "/listar", http.StatusSeeOther)
}
