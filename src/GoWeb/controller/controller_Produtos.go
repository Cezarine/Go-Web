package controller

import (
	"GoWeb/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(res http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(res, "index", model.GetProdutos())
}

func CadastroProdutos(res http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(res, "cadastraProdutos", nil)
}

func Insert(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		prod_descricao := req.FormValue("prod_descricao")
		prod_preco, err := strconv.ParseFloat(req.FormValue("prod_preco"), 64)
		if err != nil {
			log.Fatalln("Erro ao pegar preço do produto:", err.Error())
		}
		prod_quantidade, err := strconv.ParseFloat(req.FormValue("prod_quantidade"), 64)
		if err != nil {
			log.Fatalln("Erro ao pegar quantidade do produto:", err.Error())
		}
		model.CreateNewProduct(prod_descricao, prod_preco, prod_quantidade)
	}
	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

func Delete(res http.ResponseWriter, req *http.Request) {
	cod_produto, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		log.Fatalln("Error in converting cod_produto", err.Error())
	}
	model.DeleteProduct(cod_produto)

	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}

func Edit(res http.ResponseWriter, req *http.Request) {
	vCodigo := req.URL.Query().Get("id")
	vProduto := model.GetProduto(vCodigo)
	temp.ExecuteTemplate(res, "edit", vProduto)
}

func Update(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		vProduto := model.Produto{}
		var err error

		vProduto.Codigo, err = strconv.Atoi(req.FormValue("cod_produto")) // Todos esses são os "name" dos campo do html editaProdutos
		if err != nil {
			log.Fatalln("Error converting product code:", err.Error())
		}
		vProduto.Descricao = req.FormValue("prod_descricao")
		vProduto.Preco, err = strconv.ParseFloat(req.FormValue("prod_preco"), 64)
		if err != nil {
			log.Fatalln("Error converting product price")
		}
		vProduto.Quantidade, err = strconv.ParseFloat(req.FormValue("prod_quantidade"), 64)
		if err != nil {
			log.Fatalln("Error converting product quantity")
		}

		model.UpdateProduct(vProduto)
	}
	http.Redirect(res, req, "/", http.StatusMovedPermanently)
}
