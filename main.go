package main

import (
	"github.com/cezarine/API-GO-GIN/models"
	"github.com/cezarine/API-GO-GIN/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Guilherme Cezarine", CPF: "46674664804", RG: "574566544"},
		{Nome: "Ana Mariana", CPF: "1234567890", RG: "098765432"},
	}
	routes.HandRequests()
}
