package domain

type IMedicalCase interface {
	GetAll() ([]MedicalCase, error)
	GetById(id int) (MedicalCase, error)
	Create(medicalCase MedicalCase) (MedicalCase, error)
	Update(medicalCase MedicalCase) (MedicalCase, error)
	GetByMedicalCase(medicalCaseId int) ([]MedicalCase, error)
}