package models

import (
	"salao/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func RetornaProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	sql, err := db.Query("select * from produtos order by nome asc")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for sql.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int

		err = sql.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	sqlInsert, err := db.Prepare("Insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	sqlInsert.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func ExcluirProduto(id string) {
	db := db.ConectaComBancoDeDados()

	sqlDelete, err := db.Prepare("Delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	sqlDelete.Exec(id)

	defer db.Close()

}

func ConsultaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	sqlQuery, err := db.Query("Select id, nome, descricao, preco, quantidade from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for sqlQuery.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = sqlQuery.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade
	}

	defer db.Close()

	return produto

}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	sqlInsert, err := db.Prepare("Update produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	sqlInsert.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
