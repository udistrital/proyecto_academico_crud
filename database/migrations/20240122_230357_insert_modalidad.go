package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type InsertModalidad_20240122_230357 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &InsertModalidad_20240122_230357{}
	m.Created = "20240122_230357"

	migration.Register("InsertModalidad_20240122_230357", m)
}

// Run the migrations
func (m *InsertModalidad_20240122_230357) Up() {
	file, err := ioutil.ReadFile("../scripts/20240122_230357_insert_modalidad_up.sql")

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
func (m *InsertModalidad_20240122_230357) Down() {
	file, err := ioutil.ReadFile("../scripts/20240122_230357_insert_modalidad_down.sql")

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
