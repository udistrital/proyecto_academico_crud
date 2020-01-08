# proyecto_academico_crud

... .... ....


## Especificación Técnica

API de proyectos academicos, Integración con CI proyecto_academico_crud master/develop


###  Requirements
Go version >= 1.8.

### Preparation:
Para usar el API, usar el comando:

```bash
go get github.com/udistrital/proyecto_academico_crud
```
### Run

 Definir los valores de las siguientes variables de entorno:

 - `PROYECTO_ACADEMICO_API_PORT`: Puerto asignado para la ejecución del API
 - `PROYECTO_ACADEMICO_CRUD__PGUSER`: Usuario de la base de datos
 - `PROYECTO_ACADEMICO_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `PROYECTO_ACADEMICO_CRUD__PGURLS`: Host de conexión
 - `PROYECTO_ACADEMICO_CRUD__PGDB`: Nombre de la base de datos
 - `PROYECTO_ACADEMICO_CRUD__PGSCHEMA`: Esquema a utilizar en la base de datos

#### Example:

```bash
PROYECTO_ACADEMICO_API_PORT=8080 PROYECTO_ACADEMICO_CRUD__PGUSER=postgres PROYECTO_ACADEMICO_CRUD__PGPASS=**** PROYECTO_ACADEMICO_CRUD__PGURLS=localhost PROYECTO_ACADEMICO_CRUD__PGDB=proyecto_academico PROYECTO_ACADEMICO_CRUD__PGSCHEMA=proyecto_academico RUN_MODE=dev bee run -downdoc=true -gendoc=true
```
### Model DB
![image](./modelo_proyecto_academico_crud.png).

## Licencia

This file is part of proyecto_academico_crud.

proyecto_academico_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Foobar is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Foobar. If not, see https://www.gnu.org/licenses/.
