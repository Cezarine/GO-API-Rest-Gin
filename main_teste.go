package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupRotasTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestFalhador(t *testing.T) {
	t.Fatalf("Teste falhou de proposito")
}
