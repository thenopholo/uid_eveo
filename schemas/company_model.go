package schemas

import "gorm.io/gorm"

type CompanyModel struct {
	gorm.Model
	UID int
	NetsuiteID int
	TopdeskID int
	HubspotID int
	Nome string
	RazaoSocial string
	CNPJ string
	Dominio string
	NumrFuncionarios int
	ARR float64
	Setor string
	LinkedIn string
}