package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TrProyectoAcademico struct {
	ProyectoAcademicoInstitucion *ProyectoAcademicoInstitucion
	Registro                     *[]RegistroCalificadoAcreditacion
	Enfasis                      *[]ProyectoAcademicoEnfasis
	Titulaciones                 *[]Titulacion
}

type TrProyectoAcademicoPutInfoBasica struct {
	ProyectoAcademicoInstitucion *ProyectoAcademicoInstitucion
	Titulaciones                 *[]Titulacion
}

type TrProyectoAcademicoPutEnfasis struct {
	ProyectoAcademicoInstitucion *ProyectoAcademicoInstitucion
	Enfasis                      *[]ProyectoAcademicoEnfasis
}

type TrProyectoAcademicoPutRegistro struct {
	ProyectoAcademicoInstitucion *ProyectoAcademicoInstitucion
	Registro                     *[]RegistroCalificadoAcreditacion
}

// GetProyectoAcademicasById Transacción para consultar todas las proyectos academicos con toda la información de las mismas
func GetProyectoAcademicasById(id int) (v []interface{}, err error) {
	fmt.Println("entro al modelo")
	o := orm.NewOrm()
	var proyectos []*ProyectoAcademicoInstitucion

	if _, err := o.QueryTable(new(ProyectoAcademicoInstitucion)).RelatedSel().Filter("Id", id).Filter("Activo", true).All(&proyectos); err == nil {

		for _, proyecto := range proyectos {

			proyectoAcademico := proyecto
			fmt.Println("err ojoooo /n", proyecto)

			var registroProyectos []RegistroCalificadoAcreditacion
			if _, err := o.QueryTable(new(RegistroCalificadoAcreditacion)).RelatedSel().Filter("Id", id).All(&registroProyectos); err != nil {
				//fmt.Println("registro/n", registroProyectos)
				return nil, err
			}

			var enfasiproyectos []ProyectoAcademicoEnfasis
			if _, err := o.QueryTable(new(ProyectoAcademicoEnfasis)).RelatedSel().Filter("Id", id).All(&enfasiproyectos); err != nil {
				return nil, err
			}
			var titulacionproyectos []Titulacion
			if _, err := o.QueryTable(new(Titulacion)).RelatedSel().Filter("Id", id).All(&titulacionproyectos); err != nil {
				return nil, err
			}

			v = append(v, map[string]interface{}{
				"ProyectoAcademico": proyectoAcademico,
				"Registro":          &registroProyectos,
				"Enfasis":           &enfasiproyectos,
				"Titulaciones":      titulacionproyectos,
			})
		}

		return v, nil
	}
	return nil, err
}

// GetProyectoAcademicasAll Transacción para consultar todas las proyectos academicos con toda la información de las mismas
func GetProyectoAcademicasAll() (v []interface{}, err error) {
	fmt.Println("entro al modelo")
	o := orm.NewOrm()
	var proyectos []*ProyectoAcademicoInstitucion

	if _, err := o.QueryTable(new(ProyectoAcademicoInstitucion)).RelatedSel().Filter("Activo", true).All(&proyectos); err == nil {

		for _, proyecto := range proyectos {

			proyectoAcademico := proyecto
			id := proyectoAcademico.Id
			var registroProyectos []RegistroCalificadoAcreditacion
			if _, err := o.QueryTable(new(RegistroCalificadoAcreditacion)).RelatedSel().Filter("Id", id).All(&registroProyectos); err != nil {
				//fmt.Println("registro/n", registroProyectos)
				return nil, err
			}

			v = append(v, map[string]interface{}{
				"ProyectoAcademico": proyectoAcademico,
				"Registro":          registroProyectos,
			})
		}

		return v, nil
	}
	return nil, err
}

// AddTransaccionProyectoAcademica Transacción para registrar toda la información de un proyecto academico
func AddTransaccionProyectoAcademica(m *TrProyectoAcademico) (err error) {
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

		for _, v := range *m.Titulaciones {
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

// UpdateTransaccionProyectoAcademico Transacción para actualizar toda la información básicoa de un proyecto academico
func UpdateTransaccionProyectoAcademico(m *TrProyectoAcademicoPutInfoBasica) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	//Actualizar el proyecto academico
	if idProyecto, errTr := o.Update(m.ProyectoAcademicoInstitucion); errTr == nil {
		fmt.Println(idProyecto)

		for _, v := range *m.Titulaciones {
			var titulacion Titulacion
			fmt.Println("TipoTitulacionId__Id", v.TipoTitulacionId.Id, "ProyectoAcademicoInstitucionId__Id", v.ProyectoAcademicoInstitucionId.Id)
			if errTr := o.QueryTable(new(Titulacion)).RelatedSel().Filter("TipoTitulacionId__Id", v.TipoTitulacionId.Id).Filter("ProyectoAcademicoInstitucionId__Id", v.ProyectoAcademicoInstitucionId.Id).One(&titulacion); errTr == nil {
				// Si existe hace update
				v.Id = titulacion.Id
				if _, errTr = o.Update(&v, "Nombre", "Descripcion", "Activo", "NumeroOrden"); errTr != nil {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				} else {
					fmt.Println("update para titulacion", titulacion.Id)
				}
			} else {
				if errTr == orm.ErrNoRows {
					v.ProyectoAcademicoInstitucionId.Id = m.ProyectoAcademicoInstitucion.Id
					if idTitulacion, errTr := o.Insert(&v); errTr != nil {
						err = errTr
						fmt.Println(err)
						_ = o.Rollback()
						return
					} else {
						fmt.Println("insert para titulación", idTitulacion)
					}
				} else {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				}
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

// UpdateTransaccionProyectoAcademicoEnfasis Transacción para actualizar toda los enfasis de un proyecto academico
func UpdateTransaccionProyectoAcademicoEnfasis(m *TrProyectoAcademicoPutEnfasis) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	for _, v := range *m.Enfasis {
		var institucionEnfasis ProyectoAcademicoEnfasis
		fmt.Println("EnfasisId__Id", v.EnfasisId.Id, "ProyectoAcademicoInstitucionId__Id", v.ProyectoAcademicoInstitucionId.Id)
		if errTr := o.QueryTable(new(ProyectoAcademicoEnfasis)).RelatedSel().Filter("EnfasisId__Id", v.EnfasisId.Id).Filter("ProyectoAcademicoInstitucionId__Id", v.ProyectoAcademicoInstitucionId.Id).One(&institucionEnfasis); errTr == nil {
			// Si existe hace update
			v.Id = institucionEnfasis.Id
			if _, errTr = o.Update(&v, "Activo"); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			} else {
				fmt.Println("update para enfasis", institucionEnfasis.Id)
			}
		} else {
			if errTr == orm.ErrNoRows {
				v.ProyectoAcademicoInstitucionId.Id = m.ProyectoAcademicoInstitucion.Id
				if idEnfasis, errTr := o.Insert(&v); errTr != nil {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				} else {
					fmt.Println("insert para enfasis", idEnfasis)
				}
			} else {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}
	}
	_ = o.Commit()

	return
}

// UpdateTransaccionProyectoAcademicoRegistro Transacción para actualizar toda los registros de un proyecto academico
func UpdateTransaccionProyectoAcademicoRegistro(m *TrProyectoAcademicoPutRegistro) (err error) {
	o := orm.NewOrm()
	err = o.Begin()

	for _, v := range *m.Registro {
		var registro RegistroCalificadoAcreditacion
		fmt.Println("TipoRegistroId__Id", v.TipoRegistroId.Id, "ProyectoAcademicoInstitucionId__Id", v.ProyectoAcademicoInstitucionId.Id)
		if errTr := o.QueryTable(new(RegistroCalificadoAcreditacion)).RelatedSel().Filter("TipoRegistroId__Id", v.TipoRegistroId.Id).Filter("ProyectoAcademicoInstitucionId__Id", v.ProyectoAcademicoInstitucionId.Id).One(&registro); errTr == nil {
			// Si existe hace update
			v.Id = registro.Id
			if _, errTr = o.Update(&v, "NumeroActoAdministrativo", "AnoActoAdministrativoId", "FechaCreacionActoAdministrativo", "VigenciaActoAdministrativo", "VencimientoActoAdministrativo", "EnlaceActo", "Activo"); errTr != nil {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			} else {
				fmt.Println("update para registro", registro.Id)
			}
		} else {
			if errTr == orm.ErrNoRows {
				v.ProyectoAcademicoInstitucionId.Id = m.ProyectoAcademicoInstitucion.Id
				if idRegistro, errTr := o.Insert(&v); errTr != nil {
					err = errTr
					fmt.Println(err)
					_ = o.Rollback()
					return
				} else {
					fmt.Println("insert para registro", idRegistro)
				}
			} else {
				err = errTr
				fmt.Println(err)
				_ = o.Rollback()
				return
			}
		}
	}
	_ = o.Commit()

	return
}
