package main

import (
	"fmt"
	"github.com/davixcky/apiEdTeam/controllers"
	"github.com/davixcky/apiEdTeam/terminal"
	"io/ioutil"
	"log"
)

func main() {
	run()
}

func run() {
	var prompt = terminal.Prompt{}
	var actions = controllers.Actions{}

	for {
		opt, err := prompt.Draw()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(opt)
		if opt == "Salir" {
			break
		}

		if !(opt == "Crear" || opt == "Listar") {
			id, err := prompt.WaitId()
			if err != nil {
				log.Fatal(err)
			}

			resp, code, err := actions.Run(opt, id)
			if err != nil {
				log.Fatal(err)
			}

			data, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(data))
			fmt.Println(code)
			if code == 404 {
				fmt.Println("El recurso no existe 🚫")
			}
		} else {
			resp, code, err := actions.RunWithoutId(opt)
			if err != nil {
				log.Fatal(err)
			}

			data, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(data))
			fmt.Println(code)
			if code == 404 {
				fmt.Println("El recurso no existe 🚫")
			}
		}

	}
}
