package models

import "github/rafaellimasnp/aplicacaoweb/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {

	DBase := db.ConectaBD()

	selectProdutos, err := DBase.Query("SELECT * FROM produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}

	produto := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produto = append(produto, p)

	}

	defer DBase.Close()
	return produto

}

func CriarNovoProduto(nome string, descricao string, quantidade int, preco float64) {

	DBase := db.ConectaBD()

	insereProdutos, err := DBase.Prepare("insert into produtos(nome, descricao, quantidade, preco) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereProdutos.Exec(nome, descricao, quantidade, preco)

	defer DBase.Close()

}
