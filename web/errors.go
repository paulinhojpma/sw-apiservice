package web

import (
	"errors"
)

var (
	ErrorInternalServer      = errors.New("Erro desconhecido")
	ErrorFormatoRequest      = errors.New("Formato do corpo da mensagem inválido")
	ErrorPlanetaNEncontrando = errors.New("Planeta não encontrado")
)
