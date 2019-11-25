package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ProyectoAcademicoInstitucion struct {
	Id                       int             							 `orm:"column(id);pk;auto"`
	Codigo                   string          							 `orm:"column(codigo)"`
	Nombre                   string          							 `orm:"column(nombre)"`
	CodigoSnies              string          							 `orm:"column(codigo_snies)"`
	Duracion                 float64         							 `orm:"column(duracion)"`
	CorreoElectronico        string          							 `orm:"column(correo_electronico)"`
	NumeroCreditos           float64         							 `orm:"column(numero_creditos)"`
	CiclosPropedeuticos      bool            							 `orm:"column(ciclos_propedeuticos)"`
	NumeroActoAdministrativo float64         							 `orm:"column(numero_acto_administrativo)"`
	EnlaceActoAdministrativo string          							 `orm:"column(enlace_acto_administrativo)"`
	Competencias             string          							 `orm:"column(competencias)"`
	CodigoAbreviacion        string          							 `orm:"column(codigo_abreviacion);null"`
	Activo                   bool            							 `orm:"column(activo)"`
	FechaCreacion            time.Time       							 `orm:"column(fecha_creacion);type(timestamp without time zone);auto_now_add"`
	FechaModificacion        time.Time       							 `orm:"column(fecha_modificacion);type(timestamp without time zone);auto_now"`
	UnidadTiempoId           int             							 `orm:"column(unidad_tiempo_id)"`
	AnoActoAdministrativo    string          							 `orm:"column(ano_acto_administrativo)"`
	Oferta                   bool            							 `orm:"column(oferta)"`
	DependenciaId            int             							 `orm:"column(dependencia_id)"`
	AreaConocimientoId       int             							 `orm:"column(area_conocimiento_id)"`
	NucleoBaseId             int             							 `orm:"column(nucleo_base_id)"`
	MetodologiaId            *Metodologia    							 `orm:"column(metodologia_id);rel(fk)"`
	NivelFormacionId         *NivelFormacion  						 `orm:"column(nivel_formacion_id);rel(fk)"`
	FacultadId               int             							 `orm:"column(facultad_id)"`
	ProyectoPadreId          *ProyectoAcademicoInstitucion `orm:"column(proyecto_padre_id);rel(fk);null"`
}

func (t *ProyectoAcademicoInstitucion) TableName() string {
	return "proyecto_academico_institucion"
}

func init() {
	orm.RegisterModel(new(ProyectoAcademicoInstitucion))
}

// AddProyectoAcademicoInstitucion insert a new ProyectoAcademicoInstitucion into database and returns
// last inserted Id on success.
func AddProyectoAcademicoInstitucion(m *ProyectoAcademicoInstitucion) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetProyectoAcademicoInstitucionById retrieves ProyectoAcademicoInstitucion by Id. Returns error if
// Id doesn't exist
func GetProyectoAcademicoInstitucionById(id int) (v *ProyectoAcademicoInstitucion, err error) {
	o := orm.NewOrm()
	v = &ProyectoAcademicoInstitucion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllProyectoAcademicoInstitucion retrieves all ProyectoAcademicoInstitucion matches certain condition. Returns empty list if
// no records exist
func GetAllProyectoAcademicoInstitucion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ProyectoAcademicoInstitucion)).RelatedSel()
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

	var l []ProyectoAcademicoInstitucion
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

// UpdateProyectoAcademicoInstitucion updates ProyectoAcademicoInstitucion by Id and returns error if
// the record to be updated doesn't exist
func UpdateProyectoAcademicoInstitucionById(m *ProyectoAcademicoInstitucion) (err error) {
	o := orm.NewOrm()
	v := ProyectoAcademicoInstitucion{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteProyectoAcademicoInstitucion deletes ProyectoAcademicoInstitucion by Id and returns error if
// the record to be deleted doesn't exist
func DeleteProyectoAcademicoInstitucion(id int) (err error) {
	o := orm.NewOrm()
	v := ProyectoAcademicoInstitucion{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ProyectoAcademicoInstitucion{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
