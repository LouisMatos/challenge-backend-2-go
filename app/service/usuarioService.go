package service

import (
	"encoding/base64"
	"errors"
	"regexp"

	"github.com/LouisMatos/challenge-backend-2-go/app/model"
	"github.com/LouisMatos/challenge-backend-2-go/app/repository"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type UsuarioService interface {
	Save(usuario model.Usuario) (model.Usuario, error)
	Login(username string, password string) bool
}

type usuarioService struct {
	usuarioRepository repository.UsuarioRepository
}

func NewUsuarioService(repo repository.UsuarioRepository) UsuarioService {
	return &usuarioService{
		usuarioRepository: repo,
	}
}

func (service *usuarioService) Login(email string, password string) bool {

	usuario := service.usuarioRepository.GetByEmail(email)

	// comparar senhas
	pwdEncrypt := base64.StdEncoding.EncodeToString([]byte(password))

	return usuario.Password == pwdEncrypt
}

func (service *usuarioService) Save(usuario model.Usuario) (model.Usuario, error) {
	//validar email
	if (len(usuario.Email) < 5 && len(usuario.Email) > 254) || !emailRegex.MatchString(usuario.Email) {
		return usuario, errors.New("email inválido")
	}

	alreadyRegistered := service.usuarioRepository.GetByEmail(usuario.Email)

	if alreadyRegistered.Id == 0 {
		// criptografar a senha
		pwdEncrypt := base64.StdEncoding.EncodeToString([]byte(usuario.Password))
		usuario.Password = pwdEncrypt
		return service.usuarioRepository.Save(usuario), nil

	} else {
		return usuario, errors.New("Email já cadastrado")
	}

}
