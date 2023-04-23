package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome  string `json:"nome"  gorm:"not null" validate:"nonzero"`
	CPF   string `json:"cpf"   validate:"len=14, regexp=^[0-9.-x-X]*$"`
	RG    string `json:"rg"    validate:"len=12, regexp=^[0-9.-x-X]*$"`
	Ativo bool   `json:"Ativo" gorm:"default:true; not null"`
}

func ValidaDadosAlunos(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
