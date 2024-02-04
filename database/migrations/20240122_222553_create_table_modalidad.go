package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateTableModalidad_20240122_222553 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateTableModalidad_20240122_222553{}
	m.Created = "20240122_222553"

	migration.Register("CreateTableModalidad_20240122_222553", m)
}

// Run the migrations
func (m *CreateTableModalidad_20240122_222553) Up() {
	file, err := ioutil.ReadFile("../scripts/20240122_222553_create_table_modalidad_up.sql")

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
func (m *CreateTableModalidad_20240122_222553) Down() {
	file, err := ioutil.ReadFile("../scripts/20240122_222553_create_table_modalidad_down.sql")

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
