package controllers

import (
	"encoding/json"
	"fmt"

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
	//c.Mapping("GetAllById", c.GetAllById)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Put", c.Put)
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
