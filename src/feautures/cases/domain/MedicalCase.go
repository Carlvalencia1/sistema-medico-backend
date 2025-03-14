package domain

type MedicalCase struct {
	IDUsuario      int    `json:"id_usuario"`
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
	Edad           int    `json:"edad"`
	NumeroContacto string `json:"numero_contacto"`
	Descrpcion     string `json:"descrpcion"`
	Estado         string `json:"estado"`
}
