package model

import (
	"GoWeb/db"
)

type Produto struct {
	Cod        int
	Descricao  string
	Preco      float64
	Quantidade float64
}

func GetProdutos() []Produto {
	db := db.ConnectDatabase()
	produtos := []Produto{}
	p := Produto{}

	selectProducts, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	for selectProducts.Next() {
		var cod_produto int
		var prod_descricao string
		var prod_preco, prod_quantidade float64

		err = selectProducts.Scan(&cod_produto, &prod_descricao, &prod_preco, &prod_quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Cod = cod_produto
		p.Descricao = prod_descricao
		p.Preco = prod_preco
		p.Quantidade = prod_quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CreateNewProduct(pDescricao string, pPreco, pQuantidade float64) {
	db := db.ConnectDatabase()
	insertProducts, err := db.Prepare("insert into produtos (prod_descricao, prod_preco, prod_quantidade) values(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insertProducts.Exec(pDescricao, pPreco, pQuantidade)
	defer db.Close()
}
