package routes

import (
	"GoWeb/controller"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/cadastraProdutos", controller.CadastroProdutos)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
}
