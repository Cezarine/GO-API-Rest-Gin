package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome  string `json:"nome" gorm:"not null"`
	CPF   string `json:"cpf"`
	RG    string `json:"rg"`
	Ativo bool   `json:"Ativo" gorm:"default:true; not null"`
}
