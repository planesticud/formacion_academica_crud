# formacion_academica_crud
API de formación académica, Integración con CI
formacion_academica_crud master/develop
 ## Requirements
Go version >= 1.8.
 ## Preparation:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/formacion_academica_crud
 ## Run
 Definir los valores de las siguientes variables de entorno:
  - `FORMACION_ACADEMICA_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `FORMACION_ACADEMICA_CRUD__PGUSER`: Usuario de la base de datos
 - `FORMACION_ACADEMICA_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `FORMACION_ACADEMICA_CRUD__PGURLS`: Host de conexión
 - `FORMACION_ACADEMICA_CRUD__PGDB`: Nombre de la base de datos
 - `FORMACION_ACADEMICA_CRUD__SCHEMA`: Esquema a utilizar en la base de datos
 ## Example:
FORMACION_ACADEMICA_HTTP_PORT=8095 FORMACION_ACADEMICA_CRUD__PGUSER=postgres FORMACION_ACADEMICA_CRUD__PGPASS=password FORMACION_ACADEMICA_CRUD__PGURLS=localhost FORMACION_ACADEMICA_CRUD__PGDB=local FORMACION_ACADEMICA_CRUD__SCHEMA=core_new bee run
 ## Model DB
![image](https://github.com/udistrital/formacion_academica_crud/blob/dev/modelo_formacion_academica.png).
