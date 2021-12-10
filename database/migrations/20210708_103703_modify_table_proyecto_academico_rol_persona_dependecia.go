package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ModifyTableProyectoAcademicoRolPersonaDependecia_20210708_103703 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ModifyTableProyectoAcademicoRolPersonaDependecia_20210708_103703{}
	m.Created = "20210708_103703"

	migration.Register("ModifyTableProyectoAcademicoRolPersonaDependecia_20210708_103703", m)
}

// Run the migrations
func (m *ModifyTableProyectoAcademicoRolPersonaDependecia_20210708_103703) Up() {
	file, err := ioutil.ReadFile("../scripts/20210708_103703_modify_table_proyecto_academico_rol_persona_dependecia_up.sql")

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
func (m *ModifyTableProyectoAcademicoRolPersonaDependecia_20210708_103703) Down() {
	file, err := ioutil.ReadFile("../scripts/20210708_103703_modify_table_proyecto_academico_rol_persona_dependecia_down.sql")

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
