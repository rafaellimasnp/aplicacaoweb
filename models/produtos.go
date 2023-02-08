package models

import (
	"github/rafaellimasnp/aplicacaoweb/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {

	DBase := db.ConectaBD()

	selectProdutos, err := DBase.Query("SELECT * FROM produtos ORDER BY id asc")

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

func DeletaProduto(id string) {
	DBase := db.ConectaBD()

	deletarOProduto, err := DBase.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer DBase.Close()

}

func EditaProduto(id string) Produto {
	DBase := db.ConectaBD()

	produtoDoBanco, err := DBase.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Quantidade = quantidade
		produtoParaAtualizar.Preco = preco
	}

	defer DBase.Close()
	return produtoParaAtualizar

}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {

	DBase := db.ConectaBD()

	AtualizaProduto, err := DBase.Prepare("update produtos set nome=$1 , descricao=$2, preco=$3, quantidade=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)

	defer DBase.Close()
}
