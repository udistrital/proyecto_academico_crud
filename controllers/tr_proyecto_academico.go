package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/udistrital/proyecto_academico_crud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations for TrProyectoAcademico
type TrProyectoAcademicoController struct {
	beego.Controller
}

func (c *TrProyectoAcademicoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetById", c.GetById)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Put", c.Put)
}

// GetById ...
// @Title Get  By Id
// @Description get TrProyectoAcademicoController
// @Param	id		path 	string	true		"Id"
// @Success 200 {object} models.TrProyectoAcademicoController
// @Failure 404 not found resource
// @router /:id [get]
func (c *TrProyectoAcademicoController) GetById() {
	fmt.Println("entro a get")
	idProyectoStr := c.Ctx.Input.Param(":id")
	idProyecto, _ := strconv.Atoi(idProyectoStr)
	l, err := models.GetProyectoAcademicasById(idProyecto)
	if err != nil {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("404")
	} else {
		if l == nil {
			l = append(l, map[string]interface{}{})
		}
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title PostTrProyectoAcademica
// @Description create the TrProyectoAcademico
// @Param	body		body 	models.TrProyectoAcademico	true	"body for TrProyectoAcademica content"
// @Success 201 {int} models.TrProyectoAcademica
// @Failure 400 the request contains incorrect syntax
// @router / [post]
func (c *TrProyectoAcademicoController) Post() {
	fmt.Println("entro")
	var v models.TrProyectoAcademico
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fmt.Println("1")
		if err := models.AddTransaccionProyectoAcademica(&v); err == nil {
			fmt.Println("2")
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		fmt.Println("ojo")
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TrProyectoAcademico
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TrProyectoAcademico	true		"body for TrProyectoAcademico content"
// @Success 200 {object} models.TrProyectoAcademico
// @Failure 400 the request contains incorrect syntax
// @router /informacion_basica/:id [put]
func (c *TrProyectoAcademicoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var v models.TrProyectoAcademico
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.ProyectoAcademicoInstitucion.Id = id
		if err := models.UpdateTransaccionProyectoAcademica(&v); err == nil {
			c.Data["json"] = v
		} else {
			logs.Error(err)
			//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
			c.Data["system"] = err
			c.Abort("400")
		}
	} else {
		logs.Error(err)
		//c.Data["development"] = map[string]interface{}{"Code": "000", "Body": err.Error(), "Type": "error"}
		c.Data["system"] = err
		c.Abort("400")
	}
	c.ServeJSON()
}
