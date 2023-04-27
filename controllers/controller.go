package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/cezarine/API-GO-GIN/database"
	"github.com/cezarine/API-GO-GIN/models"
	"github.com/gin-gonic/gin"
)

const ativo = "true"
const inativo = "false"

func GetAlunos(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Where("ativo = ?", ativo).Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func GetAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	if c.Params.ByName("ativo") == "false" {
		database.DB.Where("ativo = ? AND id = ?", inativo, id).Find(&aluno)
	} else {
		database.DB.Where("ativo = ? AND id = ?", ativo, id).Find(&aluno)
	}
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func GetAlunoCPF(c *gin.Context) {
	var aluno models.Aluno
	vAtivo, _ := strconv.ParseBool(ativo)
	cpf := c.Param("cpf")

	if err := database.DB.Where(&models.Aluno{CPF: cpf, Ativo: vAtivo}).First(&aluno).Error; err != nil || aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno não encontrado ou inativo",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:": "Olá " + nome + " tudo bem?",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	if database.DB.Where("id = ? AND ativo = ?", id, ativo).Find(&aluno); aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno " + id + " não encontrado ou já inativo",
		})
		return
	}
	if err := database.DB.Delete(&aluno, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar aluno " + id,
			"erro":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno " + id + " deletado com sucesso",
	})
}

func AtivaInativaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	vAtivaInativa := c.Params.ByName("ativo")

	if strings.ToLower(vAtivaInativa) == "ativa" {
		if database.DB.Where("id = ? AND ativo = ?", id, inativo).Find(&aluno); aluno.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Aluno " + id + " não encontrado ou já ativo",
			})
			return
		}

		if err := database.DB.Model(&aluno).Where("id = ?", id).Update("ativo", ativo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao ativar Aluno: " + id,
				"erro":    err.Error(),
			})
			return
		}
		database.DB.Where("id = ?", id).Find(&aluno)
		c.JSON(http.StatusOK, gin.H{
			"data":  "Aluno " + id + " ativado com sucesso",
			"aluno": aluno,
		})
	} else {
		if database.DB.Where("id = ? AND ativo = ?", id, ativo).Find(&aluno); aluno.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Aluno " + id + " não encontrado ou já inativo",
			})
			return
		}

		if err := database.DB.Model(&aluno).Where("id = ?", id).Update("ativo", inativo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Erro ao inativar Aluno: " + id,
				"erro":    err.Error(),
			})
			return
		}
		database.DB.Where("id = ?", id).Find(&aluno)
		c.JSON(http.StatusOK, gin.H{
			"data": "Aluno " + id + " inativado com sucesso",
		})
		c.JSON(http.StatusOK, aluno)
	}
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	if database.DB.Where("id = ? AND ativo = ?", id, ativo).Find(&aluno); aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Aluno " + id + " não encontrado ou inativo",
		})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao alterar Aluno " + id,
			"error":   err.Error(),
		})
		return
	}

	if err := models.ValidaDadosAlunos(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := database.DB.Model(&aluno).UpdateColumns(&aluno).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao alterar Aluno " + id,
			"error":   err.Error(),
		})
		return
	}

	//	database.DB.Save(&aluno)
	c.JSON(http.StatusOK, gin.H{
		"message": "Aluno " + id + " alterado com sucesso",
		"aluno":   aluno,
	})
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	c.HTML(http.StatusOK, "index.html", gin.H{
		"mensagem": "Boas vindas",
	})
}
