package database

import (
	"wikilibras-core/src/app/models"

	"github.com/jinzhu/gorm"
)

func taskTypeSeeds(db *gorm.DB) {
	db.Create(
		models.NewTaskType(
			"Propor Sinal",
			"Tarefa onde os usuários podem adicionar um sinal",
			"propor-sinal",
		),
	)
}

func stateSeeds(db *gorm.DB) {
	db.Create(
		models.NewState(
			"Aguardando Video Referencia",
			"TODO",
			"aguardando-video-referencia",
		),
	)
	db.Create(
		models.NewState(
			"Aguardando Avaliacao Video Referencia",
			"TODO",
			"aguardando-avaliacao-video-referencia",
		),
	)
	db.Create(
		models.NewState(
			"Video Referencia Aprovado",
			"TODO",
			"video-referencia-aprovado",
		),
	)
	db.Create(
		models.NewState(
			"Video Referencia Reprovado",
			"TODO",
			"video-referencia-reprovado",
		),
	)
	db.Create(
		models.NewState(
			"Aguardando Avaliacao Animacao",
			"TODO",
			"aguardando-avaliacao-animacao",
		),
	)
	db.Create(
		models.NewState(
			"Animacao Aprovada",
			"TODO",
			"animacao-aprovada",
		),
	)
	db.Create(
		models.NewState(
			"Animacao Reprovada",
			"TODO",
			"animacao-reprovada",
		),
	)
	db.Create(
		models.NewState(
			"Aguardando Publicacao",
			"TODO",
			"aguardando-publicacao",
		),
	)
	db.Create(
		models.NewState(
			"Publicado",
			"TODO",
			"publicado",
		),
	)
}

func actionSeeds(db *gorm.DB) {
	db.Create(
		models.NewAction(
			"Informe o Video Referencia do Sinal",
			"Adicionar Video",
			"Formatos aceito .mov/.mp4",
			true,
			false,
			"adicionar-video",
		),
	)
	db.Create(
		models.NewAction(
			"Avaliar Video Referencia Positivamente Ou Negativamente",
			"Avaliar Video Referencia",
			"Verifica o manual de instruções para dicas",
			true,
			true,
			"avaliar-video",
		),
	)
	db.Create(
		models.NewAction(
			"Aprovar Video Referencia",
			"Aprovar",
			"",
			true,
			true,
			"aprovar-video",
		),
	)
	db.Create(
		models.NewAction(
			"Reprovar Video Referencia",
			"Reprovar",
			"",
			true,
			true,
			"reprovar-video",
		),
	)
	db.Create(
		models.NewAction(
			"Informe a Animacao do Sinal",
			"Adicionar Animacao",
			"Formato aceito .blend",
			true,
			false,
			"adicionar-animacao",
		),
	)
	db.Create(
		models.NewAction(
			"Aprovar Animacao",
			"Aprovar",
			"",
			true,
			true,
			"aprovar-animacao",
		),
	)
	db.Create(
		models.NewAction(
			"Reprovar Animacao",
			"Reprovar",
			"",
			true,
			true,
			"reprovar-animacao",
		),
	)
	db.Create(
		models.NewAction(
			"Gere o bundle da animação e publique o video",
			"Publicar",
			"Carregue o bundle gerado",
			true,
			true,
			"publicar-sinal",
		),
	)
}

func workflowSeeds(db *gorm.DB) {
	var taskType models.TaskType
	var statePrev models.State
	var stateNext models.State
	var action models.Action

	db.Where(models.TaskType{Slug: "propor-sinal"}).First(&taskType)
	db.Where(models.State{Slug: "aguardando-video-referencia"}).First(&statePrev)
	db.Where(models.State{Slug: "aguardando-video-referencia"}).First(&stateNext)
	db.Where(models.Action{Slug: "adicionar-video"}).First(&action)

	db.Create(
		&models.Workflow{
			TaskTypeID:  taskType.ID,
			StatePrevID: statePrev.ID,
			StateNextID: stateNext.ID,
			ActionID:    action.ID,
		},
	)
}

// RunSeeds - populate tables
func RunSeeds(db *gorm.DB) {
	stateSeeds(db)
	actionSeeds(db)
	taskTypeSeeds(db)
	workflowSeeds(db)
}
