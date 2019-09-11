package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrProyectoAcademico struct {
	ProyectoAcademicoInstitucion *ProyectoAcademicoInstitucion
	Registro                     *[]RegistroCalificadoAcreditacion
	Enfasis                      *[]ProyectoAcademicoEnfasis
}

// AddTransaccionProyectoAcademica Transacción para registrar toda la información de un proyecto academico
func AddTransaccionProyectoAcademica(m *TrProyectoAcademico) (err error) {
	fmt.Println("entro add")
	o := orm.NewOrm()
	err = o.Begin()

	if idProyecto, errTr := o.Insert(m.ProyectoAcademicoInstitucion); errTr == nil {
		fmt.Println(idProyecto)

		for _, v := range *m.Registro {
			v.ProyectoAcademicoInstitucionId.Id = int(idProyecto)
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}

		for _, v := range *m.Enfasis {
			v.ProyectoAcademicoInstitucionId.Id = int(idProyecto)
			if _, errTr = o.Insert(&v); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}
		_ = o.Commit()
	} else {
		err = errTr
		fmt.Println(err)
		_ = o.Rollback()
	}

	return
}
