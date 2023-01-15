package repository

import (
	"time"

	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"gorm.io/gorm"
)

type UsuarioRepository interface {
	Save(usuario model.Usuario) model.Usuario
	GetByEmail(email string) model.Usuario
}

type usuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(dbConnection *gorm.DB) UsuarioRepository {
	return &usuarioRepository{
		db: dbConnection,
	}
}

func (repository *usuarioRepository) Save(usuario model.Usuario) model.Usuario {

	usuario.IsActive = 1
	usuario.CreatedAt = time.Now()

	repository.db.Create(&usuario)

	return usuario
}

func (repository *usuarioRepository) GetByEmail(email string) model.Usuario {

	var usuario model.Usuario

	repository.db.Where("email = ? ", email).Find(&usuario)

	return usuario
}
