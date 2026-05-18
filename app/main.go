package main

import (
	"fmt"
	"log"
	"net/http"
	"servidorHTTP/app/handlers"
	"servidorHTTP/app/utils"
)

func main() {
	utils.ConnectToDB()

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/registrar", handlers.RegistrarHandler)
	http.HandleFunc("/listar", handlers.ListarHandler)
	http.HandleFunc("/deletar", handlers.DeletarHandler)
	http.HandleFunc("/editar", handlers.EditarPageHandler)
	http.HandleFunc("/editarSalvar", handlers.EditarHandler)

	port := "3000"
	fmt.Printf("Servidor rodando em: http://127.0.0.1:%s/\n", port)

	if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
		log.Fatal(err)
	}
}
