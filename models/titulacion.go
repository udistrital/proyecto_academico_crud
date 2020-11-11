package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Titulacion struct {
	Id                             int                           `orm:"column(id);pk;auto"`
	Nombre                         string                        `orm:"column(nombre)"`
	Descripcion                    string                        `orm:"column(descripcion);null"`
	CodigoAbreviacion              string                        `orm:"column(codigo_abreviacion);null"`
	Activo                         bool                          `orm:"column(activo)"`
	NumeroOrden                    float64                       `orm:"column(numero_orden);null"`
	FechaCreacion                  string                        `orm:"column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion              string                        `orm:"column(fecha_modificacion);type(timestamp without time zone)"`
	TipoTitulacionId               *TipoTitulacion               `orm:"column(tipo_titulacion_id);rel(fk)"`
	ProyectoAcademicoInstitucionId *ProyectoAcademicoInstitucion `orm:"column(proyecto_academico_institucion_id);rel(fk)"`
}

func (t *Titulacion) TableName() string {
	return "titulacion"
}

func init() {
	orm.RegisterModel(new(Titulacion))
}

// AddTitulacion insert a new Titulacion into database and returns
// last inserted Id on success.
func AddTitulacion(m *Titulacion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTitulacionById retrieves Titulacion by Id. Returns error if
// Id doesn't exist
func GetTitulacionById(id int) (v *Titulacion, err error) {
	o := orm.NewOrm()
	v = &Titulacion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTitulacion retrieves all Titulacion matches certain condition. Returns empty list if
// no records exist
func GetAllTitulacion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Titulacion)).RelatedSel()
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

	var l []Titulacion
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

// UpdateTitulacion updates Titulacion by Id and returns error if
// the record to be updated doesn't exist
func UpdateTitulacionById(m *Titulacion) (err error) {
	o := orm.NewOrm()
	v := Titulacion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTitulacion deletes Titulacion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTitulacion(id int) (err error) {
	o := orm.NewOrm()
	v := Titulacion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Titulacion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
