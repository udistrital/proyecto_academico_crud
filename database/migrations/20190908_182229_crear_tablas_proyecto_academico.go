package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablasProyectoAcademico_20190908_182229 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablasProyectoAcademico_20190908_182229{}
	m.Created = "20190908_182229"

	migration.Register("CrearTablasProyectoAcademico_20190908_182229", m)
}

// Run the migrations
func (m *CrearTablasProyectoAcademico_20190908_182229) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20190908_182229_crear_tablas_proyecto_academico.up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearTablasProyectoAcademico_20190908_182229) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20190908_182229_crear_tablas_proyecto_academico.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
