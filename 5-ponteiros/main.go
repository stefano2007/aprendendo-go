package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b *int = &a

	fmt.Println("Variavel a:", &a)
	fmt.Println("Variavel b:", *b)

	p := NovaPessoa("João", 25)
	// gera um panic, pois o telefone é nil
	// fmt.Println(p.Nome, p.Idade, *p.telefone)

	// Imprime os dados da pessoa
	fmt.Println(p.Nome, p.Idade, p.Telefone())
	p.AtualizarIdade(26)
	fmt.Println(p.Nome, p.Idade, p.Telefone())

	p.AtualizarTelefone("4002-8922")
	fmt.Println(p.Nome, p.Idade, p.Telefone())

	deepCopyPonteiro()
}

type Pessoa struct {
	Nome     string
	Idade    int
	telefone *string
}

func NovaPessoa(nome string, idade int) Pessoa {
	return Pessoa{
		Nome:  nome,
		Idade: idade,
	}
}

func (p Pessoa) Telefone() string {
	if p.telefone == nil {
		return ""
	}

	return *p.telefone
}

func (p *Pessoa) AtualizarIdade(idade int) {
	p.Idade = idade
}

func (p *Pessoa) AtualizarTelefone(telefone string) {
	p.telefone = &telefone
}

type Pessoa2 struct {
	Nome  *string
	Idade int
}

func deepCopyPonteiro() {

	fmt.Println("-------- INICIO DEEP COPY PONTEIRO --------")

	var nome = "Pedro"

	pessoa1 := Pessoa2{
		Nome:  &nome,
		Idade: 26,
	}

	pessoa2 := pessoa1

	fmt.Println("Nome da pessoa 1:", *pessoa1.Nome)
	fmt.Println("Nome da pessoa 2:", *pessoa2.Nome)

	pessoa2.Nome = toPointer("Davi")

	fmt.Println("Nome da pessoa 1:", *pessoa1.Nome)
	fmt.Println("Nome da pessoa 2:", *pessoa2.Nome)

	fmt.Println("--------")

	pessoa3 := Pessoa2{
		Nome:  toPointer("Leonardo"),
		Idade: 26,
	}

	pessoa4 := deepCopy(pessoa3)

	*pessoa4.Nome = "Lucas"
	fmt.Println("Nome da pessoa 3:", *pessoa3.Nome)
	fmt.Println("Nome da pessoa 4:", *pessoa4.Nome)

	fmt.Println("--------")
	listaDePessoas1 := []Pessoa2{pessoa1, pessoa2, pessoa3, pessoa4}
	listaDePessoas2 := deepCopyList(listaDePessoas1)

	fmt.Println("Lista 1", listaDePessoas1)
	fmt.Println("Lista 2", listaDePessoas2)
}

func toPointer(s string) *string {
	return &s
}

func deepCopy(origem Pessoa2) Pessoa2 {
	var destino Pessoa2

	destino.Idade = origem.Idade
	destino.Nome = toPointer(*origem.Nome)

	return destino
}

func deepCopyList(origem []Pessoa2) []Pessoa2 {
	var destino = make([]Pessoa2, len(origem))

	for i, pessoa := range origem {
		destino[i] = deepCopy(pessoa)
	}

	return destino
}
