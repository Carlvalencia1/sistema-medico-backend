package domain

type IPatients interface {
	GetAll() ([]Patients, error)
	GetById(id int) (Patients, error)
	Create(patient Patients) (Patients, error)
	Update(patient Patients) (Patients, error)
	Delete(id int) error
	GetByPatient(patientId int) ([]Patients, error)
}