package filtros

import (
	"fmt"
)

type UsuarioFiltro struct {
	Id uint
}

func (uf *UsuarioFiltro) Validate() (erro error) {

	if uf.Id == 0 {
		return fmt.Errorf(ERROR_VALIDATE_ID)
	}
	return
}
