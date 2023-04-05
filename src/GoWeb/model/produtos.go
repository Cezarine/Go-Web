package model

import (
	"GoWeb/db"
	"log"
)

type Produto struct {
	Codigo     int
	Descricao  string
	Preco      float64
	Quantidade float64
}

func GetProdutos() []Produto {
	db := db.ConnectDatabase()
	produtos := []Produto{}
	p := Produto{}

	selectProducts, err := db.Query("SELECT * FROM produtos WHERE prod_inativo = 0 ORDER BY cod_produto ASC")
	if err != nil {
		panic(err.Error())
	}

	for selectProducts.Next() {
		var cod_produto, prod_inativo int
		var prod_descricao string
		var prod_preco, prod_quantidade float64

		err = selectProducts.Scan(&cod_produto, &prod_descricao, &prod_preco, &prod_quantidade, &prod_inativo)
		if err != nil {
			log.Fatalln("Erro ao pegar os produtos ativos:", err.Error())
		}

		p.Codigo = cod_produto
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
		log.Fatalln("Erro ao inserir produto novo", err.Error())
	}
	insertProducts.Exec(pDescricao, pPreco, pQuantidade)
	defer db.Close()
}

func DeleteProduct(pCodProduto int) {
	db := db.ConnectDatabase()

	delete, err := db.Prepare("UPDATE produtos SET prod_inativo = 1 WHERE cod_produto = ?")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(pCodProduto)
	defer db.Close()
}

func GetProduto(pCodProduto string) Produto {
	var vCod_produto, vProd_inativo int
	var vProd_descricao string
	var vProd_preco, vProd_quantidade float64
	vProduct := Produto{}
	db := db.ConnectDatabase()

	vSelectProduct, err := db.Query("SELECT * FROM produtos WHERE cod_produto = ?", pCodProduto)
	if err != nil {
		log.Fatalln("Product selection error:", err.Error())
	}
	defer db.Close()

	for vSelectProduct.Next() {
		err := vSelectProduct.Scan(&vCod_produto, &vProd_descricao, &vProd_preco, &vProd_quantidade, &vProd_inativo)
		if err != nil {
			log.Fatalln("Error fetching SELECT fields", err.Error())
		}
		vProduct.Codigo = vCod_produto
		vProduct.Descricao = vProd_descricao
		vProduct.Preco = vProd_preco
		vProduct.Quantidade = vProd_quantidade

	}
	return vProduct
}

func UpdateProduct(pProduto Produto) {
	db := db.ConnectDatabase()
	vUpdateProduct, err := db.Prepare("UPDATE produtos SET prod_descricao=?, prod_preco=?, prod_quantidade=? WHERE cod_produto=?")
	if err != err {
		log.Fatalln("Error updating product", pProduto.Descricao, ":", err.Error())
	}
	vUpdateProduct.Exec(pProduto.Descricao, pProduto.Preco, pProduto.Quantidade, pProduto.Codigo)
	defer db.Close()
}
