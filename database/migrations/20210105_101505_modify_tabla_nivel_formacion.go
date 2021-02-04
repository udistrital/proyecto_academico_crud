package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type ModifyTablaNivelFormacion_20210105_101505 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModifyTablaNivelFormacion_20210105_101505{}
	m.Created = "20210105_101505"

	migration.Register("ModifyTablaNivelFormacion_20210105_101505", m)
}

// Run the migrations
func (m *ModifyTablaNivelFormacion_20210105_101505) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210105_152820_agregar_campos_nivel_formacion_padre.up.sql")

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
func (m *ModifyTablaNivelFormacion_20210105_101505) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210105_152820_agregar_campos_nivel_formacion_padre.down.sql")

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
