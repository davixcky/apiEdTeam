package terminal

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

type Prompt struct{}

// Dibuja el menu, es decir, las opciones que contendra
func (*Prompt) Draw() (option string, err error) {
	// Opciones del menu
	prompt := promptui.Select{
		Label: "Seleccione una opcion",
		Items: []string{"Crear", "Consultar", "Listar", "Modificar", "Eliminar", "Salir"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

// Espera la digitalizacion del id y lo verifica internamente
func (*Prompt) WaitId() (id int, err error) {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Numero invalido")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Numero",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0, err
	}

	id, err = strconv.Atoi(result)
	if err != nil {
		fmt.Printf("Strconv failed %v\n", err)
		return 0, err
	}

	return id, nil
}

