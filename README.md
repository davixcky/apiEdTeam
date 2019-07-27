# Ed Team - Reto Backend!

Este proyecto es la solución a un problema, realizado por la emprea Ed Team, con la finalidad de obtener una vacante en el area del **backend**.

La aplicacion consiste en un CLI, el cual debe contar un menu con 5 opciones principales, crear, editar, eliminar, consultar y listar, consumiendo la siguiente [api](https://jsonplaceholder.typicode.com/photos).


# Esquema de archivos

El proyecto cuenta con 4 carpetas principales

 - Controllers
 - Photos
 - Slack
 - Terminal

## Controllers

Controla y maneja todas las acciones e interacciones de la pantalla con los recursos, es decir, es el conductor de los actuadores.

## Photos

Llama directamente a la api ofrecida por [jsonplaceholder](https://jsonplaceholder.typicode.com/photos).

## Slack

Accede directamente a slack mediante llamadas por webhook.

## Terminal

Renderiza y controla las acciones del menú y la terminal .

# Descarga
Puedes descargar el codigo mediante el comando `https://github.com/davixcky/apiEdTeam.git` o simplemente descargando el `.zip`

# Ejecutar

 ### 1. Iniciar en el grupo de Slack
 Para entrar en el grupo, debe iniciar a traves del siguiente [enlace](https://join.slack.com/t/pruebaedteam/shared_invite/enQtNzA4NTIyNDExNTU5LTdmZjExYjI1ZTNiNjQzMmM1N2JiY2U2MDNmNjljMGYwNGYzYzczMDNlYzliMGFiNjY4YTI5MTM5ZTY3MDE2YjM).

### 2. Ejecutar la aplicacion
Puede ejecutar la aplicación compilando con el comando `go build` y posterior a ello ejecutar según su sistema operativo o simplemente ejecutarla con el siguiente comando `go run *.go`.



