package controllers

import (
	"github.com/cezarine/API-GO-GIN/models"
	"github.com/gin-gonic/gin"
)

func GetAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Olá " + nome + " tudo bem?",
	})
}
