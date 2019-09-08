package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type RegistroCalificadoAcreditaciÛn struct {
	Id                              int                           `orm:"column(id);pk"`
	NumeroActoAdministrativo        float64                       `orm:"column(numero_acto_administrativo)"`
	AnoActoAdministrativoId         int                           `orm:"column(ano_acto_administrativo_id)"`
	FechaCreacionActoAdministrativo time.Time                     `orm:"column(fecha_creacion_acto_administrativo);type(timestamp without time zone)"`
	VigenciaActoAdministrativo      string                        `orm:"column(vigencia_acto_administrativo)"`
	VencimientoActoAdministrativo   time.Time                     `orm:"column(vencimiento_acto_administrativo);type(timestamp without time zone)"`
	EnlaceActo                      string                        `orm:"column(enlace_acto)"`
	Activo                          bool                          `orm:"column(activo)"`
	FechaCreacion                   time.Time                     `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion               time.Time                     `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	ProyectoAcademicoInstitucionId  *ProyectoAcademicoInstitucion `orm:"column(proyecto_academico_institucion_id);rel(fk)"`
	TipoRegistroId                  *TipoRegistro                 `orm:"column(tipo_registro_id);rel(fk)"`
}

func (t *RegistroCalificadoAcreditaciÛn) TableName() string {
	return "registro_calificado_acreditaciÛn"
}

func init() {
	orm.RegisterModel(new(RegistroCalificadoAcreditaciÛn))
}

// AddRegistroCalificadoAcreditaciÛn insert a new RegistroCalificadoAcreditaciÛn into database and returns
// last inserted Id on success.
func AddRegistroCalificadoAcreditaciÛn(m *RegistroCalificadoAcreditaciÛn) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRegistroCalificadoAcreditaciÛnById retrieves RegistroCalificadoAcreditaciÛn by Id. Returns error if
// Id doesn't exist
func GetRegistroCalificadoAcreditaciÛnById(id int) (v *RegistroCalificadoAcreditaciÛn, err error) {
	o := orm.NewOrm()
	v = &RegistroCalificadoAcreditaciÛn{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRegistroCalificadoAcreditaciÛn retrieves all RegistroCalificadoAcreditaciÛn matches certain condition. Returns empty list if
// no records exist
func GetAllRegistroCalificadoAcreditaciÛn(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RegistroCalificadoAcreditaciÛn))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []RegistroCalificadoAcreditaciÛn
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRegistroCalificadoAcreditaciÛn updates RegistroCalificadoAcreditaciÛn by Id and returns error if
// the record to be updated doesn't exist
func UpdateRegistroCalificadoAcreditaciÛnById(m *RegistroCalificadoAcreditaciÛn) (err error) {
	o := orm.NewOrm()
	v := RegistroCalificadoAcreditaciÛn{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRegistroCalificadoAcreditaciÛn deletes RegistroCalificadoAcreditaciÛn by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRegistroCalificadoAcreditaciÛn(id int) (err error) {
	o := orm.NewOrm()
	v := RegistroCalificadoAcreditaciÛn{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&RegistroCalificadoAcreditaciÛn{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
