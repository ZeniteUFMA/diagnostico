package models

import "time"

type SurveyResponse struct {
	ID uint `gorm:"primaryKey" json:"id"`

	// Dados básicos
	Idade      string `json:"idade"`
	Curso      string `json:"curso"`
	Semestre   string `json:"semestre"`
	Faculdade  string `json:"faculdade"`

	// Momento profissional
	SituacaoProfissional string `json:"situacao_profissional"`
	ParticipouProcessos  string `json:"participou_processos"`
	CurriculoAtualizado  string `json:"curriculo_atualizado"`
	LinkedinProfissional string `json:"linkedin_profissional"`

	// Clareza de carreira (1-5)
	AreaClara                 int `json:"area_clara"`
	ObjetivosDefinidos        int `json:"objetivos_definidos"`
	CompetenciasDesenvolver   int `json:"competencias_desenvolver"`
	SegurancaFuturo           int `json:"seguranca_futuro"`
	PlanoParaObjetivos        int `json:"plano_para_objetivos"`
	EntendeMercado            int `json:"entende_mercado"`

	// Dificuldades
	MaiorDificuldade string `json:"maior_dificuldade"`
	PreocupacaoAtual string `json:"preocupacao_atual"`

	// Interesses
	AreasInteresse string `json:"areas_interesse"`

	// Objetivo atual
	ObjetivoAtual string `json:"objetivo_atual"`

	// Competências (1-5)
	Comunicacao           int `json:"comunicacao"`
	Organizacao           int `json:"organizacao"`
	Lideranca             int `json:"lideranca"`
	InteligenciaEmocional int `json:"inteligencia_emocional"`
	TrabalhoEquipe        int `json:"trabalho_equipe"`
	Planejamento          int `json:"planejamento"`
	CapacidadeAnalitica   int `json:"capacidade_analitica"`
	Networking            int `json:"networking"`

	// Desenvolvimento profissional
	InvesteDesenvolvimento      string `json:"investe_desenvolvimento"`
	CursosExtracurriculares     string `json:"cursos_extracurriculares"`
	TempoSemanalDesenvolvimento string `json:"tempo_semanal_desenvolvimento"`

	CreatedAt time.Time `json:"created_at"`
}
