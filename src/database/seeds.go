package database

import (
	"wikilibras-core/src/app/models"

	"github.com/jinzhu/gorm"
)

func taskTypeSeeds(db *gorm.DB) {
	db.Create(&models.TaskType{
		Name:        "Propor Sinal",
		Description: "Tarefa onde os usuários podem adicionar um sinal",
		Slug:        "propor-sinal",
	},
	)
}

func stateSeeds(db *gorm.DB) {
	db.Create(
		&models.State{
			Name:        "Aguardando Video Referencia",
			Description: "TODO",
			Slug:        "aguardando-video-referencia",
		},
	)
	db.Create(
		&models.State{
			Name:        "Aguardando Avaliacao Video Referencia",
			Description: "TODO",
			Slug:        "aguardando-avaliacao-video-referencia",
		},
	)
	db.Create(
		&models.State{
			Name:        "Video Referencia Aprovado",
			Description: "TODO",
			Slug:        "video-referencia-aprovado",
		},
	)
	db.Create(
		&models.State{
			Name:        "Video Referencia Reprovado",
			Description: "TODO",
			Slug:        "video-referencia-reprovado",
		},
	)
	db.Create(
		&models.State{
			Name:        "Aguardando Avaliacao Animacao",
			Description: "TODO",
			Slug:        "aguardando-avaliacao-animacao",
		},
	)
	db.Create(
		&models.State{
			Name:        "Animacao Aprovada",
			Description: "TODO",
			Slug:        "animacao-aprovada",
		},
	)
	db.Create(
		&models.State{
			Name:        "Animacao Reprovada",
			Description: "TODO",
			Slug:        "animacao-reprovada",
		},
	)
	db.Create(
		&models.State{
			Name:        "Aguardando Publicacao",
			Description: "TODO",
			Slug:        "aguardando-publicacao",
		},
	)
	db.Create(
		&models.State{
			Name:        "Publicado",
			Description: "TODO",
			Slug:        "publicado",
		},
	)
}

func actionSeeds(db *gorm.DB) {
	db.Create(
		&models.Action{
			Description: "Informe o Video Referencia do Sinal",
			Caption:     "Adicionar Video",
			Help:        "Formatos aceito .mov/.mp4",
			Slug:        "adicionar-video",
		},
	)
	db.Create(
		&models.Action{
			Description: "Aprovar Video Referencia",
			Caption:     "Aprovar",
			Slug:        "aprovar-video",
		},
	)
	db.Create(
		&models.Action{
			Description: "Reprovar Video Referencia",
			Caption:     "Reprovar",
			Slug:        "reprovar-video",
		},
	)
	db.Create(
		&models.Action{
			Description: "Informe a Animacao do Sinal",
			Caption:     "Adicionar Animacao",
			Help:        "Formato aceito .blend",
			Slug:        "adicionar-animacao",
		},
	)
	db.Create(
		&models.Action{
			Description: "Aprovar Animacao",
			Caption:     "Aprovar",
			Slug:        "aprovar-animacao",
		},
	)
	db.Create(
		&models.Action{
			Description: "Reprovar Animacao",
			Caption:     "Reprovar",
			Slug:        "reprovar-animacao",
		},
	)
	db.Create(
		&models.Action{
			Description: "Gere o bundle da animação e publique o video",
			Caption:     "Publicar",
			Help:        "Carregue o bundle gerado",
			Slug:        "publicar-sinal",
		},
	)
}

func workflowSeeds(db *gorm.DB) {
	var taskType models.TaskType
	var statePrev models.State
	var stateNext models.State
	var action models.Action

	db.Where(models.TaskType{Slug: "propor-sinal"}).First(&taskType)
	db.Where(models.State{Slug: "aguardando-video-referencia"}).First(&statePrev)
	db.Where(models.State{Slug: "aguardando-avaliacao-video-referencia"}).First(&stateNext)
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

func orientationSeeds(db *gorm.DB) {
	db.Create(
		&models.Orientation{
			Name:        "Positiva",
			Description: "TODO",
			Slug:        "positiva",
		},
	)
	db.Create(
		&models.Orientation{
			Name:        "Negativa",
			Description: "TODO",
			Slug:        "negativa",
		},
	)
	db.Create(
		&models.Orientation{
			Name:        "Indefinida",
			Description: "TODO",
			Slug:        "indefinida",
		},
	)
}

// RunSeeds - populate tables
func RunSeeds(db *gorm.DB) {
	stateSeeds(db)
	actionSeeds(db)
	taskTypeSeeds(db)
	workflowSeeds(db)
	orientationSeeds(db)
}
