package main

import (
	"fmt"
	"github.com/astaxie/beego/migration"
	"io/ioutil"
	"strings"
)

// DO NOT MODIFY
type InsertsTablaNivelFormacion_20210105_103620 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertsTablaNivelFormacion_20210105_103620{}
	m.Created = "20210105_103620"

	migration.Register("InsertsTablaNivelFormacion_20210105_103620", m)
}

// Run the migrations
func (m *InsertsTablaNivelFormacion_20210105_103620) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210105_182229_inserts_tabla_nivel_formacion.up.sql")

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
func (m *InsertsTablaNivelFormacion_20210105_103620) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20190909_182229_inserts_parametricas.down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";\n")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}
}
