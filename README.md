# proyecto_academico_crud

Api CRUD para el manejo de la información de los proyectos académicos de la Universidad Distrital

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)



### Variables de Entorno
```shell
PROYECTO_ACADEMICO_CRUD_HTTP_PORT: [Puerto asignado para la ejecución del API]
PROYECTO_ACADEMICO_CRUD_PGUSER: [Usuario de la base de datos]
PROYECTO_ACADEMICO_CRUD_PGPASS: [Clave del usuario para la conexión a la base de datos]
PROYECTO_ACADEMICO_CRUD_PGHOST: [Host de conexión]
PROYECTO_ACADEMICO_CRUD_PGPORT: [Puerto de conexión a la base de datos]
PROYECTO_ACADEMICO_CRUD_PGDB: [Nombre de la base de datos]
PROYECTO_ACADEMICO_CRUD_PGSCHEMA: [Esquema a utilizar en la base de datos]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con PROYECTO_ACADEMICO_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/proyecto_academico_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/proyecto_academico_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
PROYECTO_ACADEMICO_CRUD_HTTP_PORT=8080 PROYECTO_ACADEMICO_CRUD_PGHOST=127.0.0.1:27017 PROYECTO_ACADEMICO_CRUD_SOME_VARIABLE=some_value bee run
```
### Ejecución Dockerfile
```shell
# docker build --tag=proyecto_academico_crud . --no-cache
# docker run -p 80:80 proyecto_academico_crud
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/proyecto_academico_crud

#2. Moverse a la carpeta del repositorio
cd proyecto_academico_crud

#3. Crear un fichero con el nombre **custom.env**
# En windows ejecutar:* ` ni custom.env`
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas

Pruebas unitarias
```shell
# En Proceso
```
## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/proyecto_academico_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/proyecto_academico_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/proyecto_academico_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/proyecto_academico_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/proyecto_academico_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/proyecto_academico_crud) |

## Modelo de Datos
[Modelo de Datos API CRUD proyecto_academico_crud](https://user-images.githubusercontent.com/14035745/93823393-5e0c2600-fc27-11ea-8e7f-7373cf931631.png)

## Licencia

This file is part of proyecto_academico_crud.

proyecto_academico_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

proyecto_academico_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with proyecto_academico_crud. If not, see https://www.gnu.org/licenses/.
