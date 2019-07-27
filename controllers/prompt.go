package controllers

import (
	"errors"
	"fmt"
	"github.com/davixcky/apiEdTeam/photos"
	"github.com/davixcky/apiEdTeam/slack"
	"net/http"
	"strings"
)

type Actions struct{}

var photo = api.Photos{}

// Actua sobre la opcion de crear
func (*Actions) create() (resp *http.Response, code int, err error) {
	resp, code, err = photo.Post()
	if err != nil {
		return &http.Response{}, 0, err
	}

	// Envia la notificacion a slack
	var controlSlack = slack.Slack{}
	if code == 201 {
		err = controlSlack.Created("", 5001)
		if err != nil {
			return &http.Response{}, 0, err
		}
	}

	return resp, code, nil
}

// Actua sobre la opcion de consultar
func (*Actions) getOne(id int) (resp *http.Response, code int, err error) {
	fmt.Println("Cargando...")
	return photo.GetOne(id)
}

// Actua sobre la opcion de listar
func (*Actions) getAll() (resp *http.Response, code int, err error) {
	fmt.Println("Cargando...")
	return photo.GetAll()
}

// Actua sobre la opcion de eliminar
func (*Actions) delete(id int) (resp *http.Response, code int, err error) {
	fmt.Println("Cargando...")
	resp, code, err = photo.Delete(id)
	if err != nil {
		return &http.Response{}, 0, err
	}

	// Envia la notificacion a slack
	var controlSlack = slack.Slack{}
	if code == 200 {
		err = controlSlack.Deleted("", id)
		if err != nil {
			return &http.Response{}, 0, err
		}
	}

	return resp, code, nil
}

// Actua sobre la opcion de actualizar
func (*Actions) update(id int) (resp *http.Response, code int, err error) {
	fmt.Println("Cargando...")
	return photo.Update(id)
}

// Gestiona los actuadores segun su accion
func (h *Actions) Run(option string, id int) (resp *http.Response, code int, err error) {
	option = strings.ToLower(option)
	switch option {
	case "crear":
		return h.create()
	case "consultar":
		return h.getOne(id)
	case "listar":
		return h.getAll()
	case "eliminar":
		return h.delete(id)
	case "modificar":
		return h.update(id)
	default:
		return nil, 0, errors.New(fmt.Sprintf("action <%s> doesn't exists", option))
	}

}

// Gestiona los actuadores segun su accion sin necesidad de introducir un Id
func (h *Actions) RunWithoutId(option string) (resp *http.Response, code int, err error) {
	return h.Run(option, 0)
}
