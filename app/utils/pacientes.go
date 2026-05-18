package utils

import (
	"log"
	"time"
)

type Paciente struct {
	ID               int
	NomeCompleto     string
	CPF              string
	DataNascimento   string
	Telefone         string
	Email            string
	DataInternacao   string
	HoraInternacao   string
	MotivoInternacao string
}

func InsertPaciente(nomeCompleto, cpf, dataNascimento, telefone, email, dataInternacao, horaInternacao, motivoInternacao string) error {
	dataHoraCompleta := dataInternacao + " " + horaInternacao + ":00"

	query := `INSERT INTO pacientes (nome_completo, cpf, data_nascimento, telefone, email, data_internacao, motivo_internacao) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := DB.Exec(query, nomeCompleto, cpf, dataNascimento, telefone, email, dataHoraCompleta, motivoInternacao)
	if err != nil {
		log.Printf("Erro ao inserir paciente no banco de dados: %v", err)
		return err
	}
	log.Println("Paciente inserido com sucesso!")
	return nil
}

func GetAllPacientes() ([]Paciente, error) {
	query := `SELECT id, nome_completo, cpf, data_nascimento, telefone, email, data_internacao, motivo_internacao FROM pacientes ORDER BY created_at DESC`
	rows, err := DB.Query(query)
	if err != nil {
		log.Printf("Erro ao buscar pacientes no banco de dados: %v", err)
		return nil, err
	}
	defer rows.Close()

	var pacientes []Paciente
	for rows.Next() {
		var paciente Paciente
		var dataNascimentoTime time.Time
		var dataInternacaoTime time.Time

		err := rows.Scan(&paciente.ID, &paciente.NomeCompleto, &paciente.CPF, &dataNascimentoTime, &paciente.Telefone, &paciente.Email, &dataInternacaoTime, &paciente.MotivoInternacao)
		if err != nil {
			log.Printf("Erro ao escanear paciente: %v", err)
			return nil, err
		}

		paciente.DataNascimento = dataNascimentoTime.Format("02/01/2006")
		paciente.DataInternacao = dataInternacaoTime.Format("02/01/2006")
		paciente.HoraInternacao = dataInternacaoTime.Format("15:04")

		pacientes = append(pacientes, paciente)
	}

	return pacientes, nil
}

func GetPacienteByID(id int) (*Paciente, error) {
	query := `SELECT id, nome_completo, cpf, data_nascimento, telefone, email, data_internacao, motivo_internacao FROM pacientes WHERE id = $1`
	var paciente Paciente
	var dataNascimentoTime time.Time
	var dataInternacaoTime time.Time

	err := DB.QueryRow(query, id).Scan(&paciente.ID, &paciente.NomeCompleto, &paciente.CPF, &dataNascimentoTime, &paciente.Telefone, &paciente.Email, &dataInternacaoTime, &paciente.MotivoInternacao)
	if err != nil {
		log.Printf("Erro ao buscar paciente no banco de dados: %v", err)
		return nil, err
	}

	paciente.DataNascimento = dataNascimentoTime.Format("02/01/2006")
	paciente.DataInternacao = dataInternacaoTime.Format("02/01/2006")
	paciente.HoraInternacao = dataInternacaoTime.Format("15:04")

	return &paciente, nil
}

func UpdatePaciente(id int, nomeCompleto, cpf, dataNascimento, telefone, email, dataInternacao, horaInternacao, motivoInternacao string) error {
	dataHoraCompleta := dataInternacao + " " + horaInternacao + ":00"

	query := `UPDATE pacientes SET nome_completo = $1, cpf = $2, data_nascimento = $3, telefone = $4, email = $5, data_internacao = $6, motivo_internacao = $7 WHERE id = $8`

	_, err := DB.Exec(query, nomeCompleto, cpf, dataNascimento, telefone, email, dataHoraCompleta, motivoInternacao, id)
	if err != nil {
		log.Printf("Erro ao atualizar paciente no banco de dados: %v", err)
		return err
	}
	log.Println("Paciente atualizado com sucesso!")
	return nil
}

func DeletePaciente(id int) error {
	query := `DELETE FROM pacientes WHERE id = $1`
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Printf("Erro ao apagar paciente do banco de dados: %v", err)
		return err
	}
	log.Println("Paciente apagado com sucesso!")
	return nil
}
