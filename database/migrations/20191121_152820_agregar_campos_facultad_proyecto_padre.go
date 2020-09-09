package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarCamposFacultadProyectoPadre_20191121_152820 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarCamposFacultadProyectoPadre_20191121_152820{}
	m.Created = "20191121_152820"

	migration.Register("AgregarCamposFacultadProyectoPadre_20191121_152820", m)
}

// Run the migrations
func (m *AgregarCamposFacultadProyectoPadre_20191121_152820) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20191121_152820_agregar_campos_facultad_proyecto_padre.up.sql")

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
func (m *AgregarCamposFacultadProyectoPadre_20191121_152820) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20191121_152820_agregar_campos_facultad_proyecto_padre.down.sql")

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
