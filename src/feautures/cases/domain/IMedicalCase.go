package domain

type IMedicalCase interface {
	GetAll() ([]MedicalCase, error)
	GetById(id int) (MedicalCase, error)
	Create(medicalCase MedicalCase) (MedicalCase, error)
	Update(medicalCase MedicalCase) (MedicalCase, error)
	Delete(id int) error
	GetByMedicalCase(medicalCaseId int) ([]MedicalCase, error)
}