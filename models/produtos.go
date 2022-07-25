package models

import "go-produtos-alura/db"

type Produto struct {
	Id int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		p.Id = id

		produtos = append(produtos, p)
	}

	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int){
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
}

func DeletaProduto(id string){
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	deletarProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)
}