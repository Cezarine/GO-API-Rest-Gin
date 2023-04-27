package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/cezarine/API-GO-GIN/controllers"
	"github.com/cezarine/API-GO-GIN/database"
	"github.com/cezarine/API-GO-GIN/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const StatusOK = http.StatusOK

var ID int
var Nome string
var CPF string
var RG string

func SetupRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "NOME DO ALUNO TESTE",
		CPF:  "123.456.789-00",
		RG:   "12.345.678-X",
	}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
	Nome = string(aluno.Nome)
	CPF = string(aluno.CPF)
	RG = string(aluno.RG)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeSaudacao(t *testing.T) {
	r := SetupRotasTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Guilherme", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	assert.Equal(t, StatusOK, res.Code, fmt.Sprintf("Status Code Error: Status expected %d, status received %d", StatusOK, res.Code))

	mockDaResposta := `{"API diz:":"Olá Guilherme tudo bem?"}` //Mock significa mocar, pegar/guardar
	resBody, _ := io.ReadAll(res.Body)
	assert.Equal(t, mockDaResposta, string(resBody), fmt.Sprintf("Body Error: Body expected %s, body received %s", mockDaResposta, string(resBody)))
}

func TestListandoAlunosHandler(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()

	r := SetupRotasTeste()
	r.GET("/alunos", controllers.GetAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	assert.Equal(t, StatusOK, res.Code, fmt.Sprintf("Status Code Error: Status expected %d, status received %d", StatusOK, res.Code))

	defer DeletaAlunoMock()
}

func TestBuscaAlunoCPFHandler(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()

	r := SetupRotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.GetAlunoCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/123.456.789-00", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	assert.Equal(t, StatusOK, res.Code, fmt.Sprintf("Status Code Error: Status expected %d, status received %d", StatusOK, res.Code))

	defer DeletaAlunoMock()
}

func TestBuscaAlunoID(t *testing.T) {
	var alunoMock models.Aluno
	database.ConectaBanco()
	CriaAlunoMock()

	r := SetupRotasTeste()
	r.GET("/alunos/:id", controllers.GetAluno)

	pathBusca := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", pathBusca, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)
	json.Unmarshal(res.Body.Bytes(), &alunoMock)
	assert.Equal(t, StatusOK, res.Code, fmt.Sprintf("Nome Error: Nome expected %d, nome received %d", StatusOK, res.Code))
	assert.Equal(t, Nome, alunoMock.Nome, fmt.Sprintf("Nome Error: Nome expected %s, nome received %s", Nome, alunoMock.Nome))
	assert.Equal(t, CPF, alunoMock.CPF, fmt.Sprintf("Nome Error: Nome expected %s, nome received %s", CPF, alunoMock.CPF))
	assert.Equal(t, RG, alunoMock.RG, fmt.Sprintf("Nome Error: Nome expected %s, nome received %s", RG, alunoMock.RG))

	defer DeletaAlunoMock()
}

func TestDeleteAluno(t *testing.T) {
	database.ConectaBanco()
	CriaAlunoMock()

	r := SetupRotasTeste()
	r.DELETE("/alunos/deleta/:id", controllers.DeletaAluno)

	pathBusca := "/alunos/deleta/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", pathBusca, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, StatusOK, res.Code, fmt.Sprintf("Nome Error: Nome expected %d, nome received %d", StatusOK, res.Code))
}

func TestAtualizaAluno(t *testing.T) {
	var alunoMockAtualizado models.Aluno
	vAluno := models.Aluno{Nome: "NOME DO ALUNO TESTE ATUALIZADO", CPF: "000.456.789-00", RG: "00.345.678-X"}
	fmt.Print(vAluno.ID)
	AlunoJson, _ := json.Marshal(vAluno)

	database.ConectaBanco()
	CriaAlunoMock()

	r := SetupRotasTeste()
	r.PUT("/alunos/:id", controllers.EditaAluno)
	pathEditar := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("PUT", pathEditar, bytes.NewBuffer(AlunoJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	json.Unmarshal(res.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, StatusOK, res.Code, fmt.Sprintf("Status code Error: Status code expected %d, Status code received %d", StatusOK, res.Code))
	assert.Equal(t, vAluno.Nome, alunoMockAtualizado.Nome, fmt.Sprintf("Nome Error: Nome expected %s, nome received %s", vAluno.Nome, alunoMockAtualizado.Nome))
	assert.Equal(t, vAluno.CPF, alunoMockAtualizado.CPF, fmt.Sprintf("CPF Error: CPF expected %s, CPF received %s", vAluno.CPF, alunoMockAtualizado.CPF))
	assert.Equal(t, vAluno.RG, alunoMockAtualizado.RG, fmt.Sprintf("RG Error: RG expected %s, RG received %s", vAluno.RG, alunoMockAtualizado.RG))

	defer DeletaAlunoMock()
}
