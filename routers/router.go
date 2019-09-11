// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/proyecto_academico_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/metodologia",
			beego.NSInclude(
				&controllers.MetodologiaController{},
			),
		),

		beego.NSNamespace("/nivel_formacion",
			beego.NSInclude(
				&controllers.NivelFormacionController{},
			),
		),

		beego.NSNamespace("/enfasis",
			beego.NSInclude(
				&controllers.EnfasisController{},
			),
		),

		beego.NSNamespace("/proyecto_academico_enfasis",
			beego.NSInclude(
				&controllers.ProyectoAcademicoEnfasisController{},
			),
		),

		beego.NSNamespace("/tipo_registro",
			beego.NSInclude(
				&controllers.TipoRegistroController{},
			),
		),

		beego.NSNamespace("/registro_calificado_acreditacion",
			beego.NSInclude(
				&controllers.RegistroCalificadoAcreditacionController{},
			),
		),

		beego.NSNamespace("/tipo_titulacion",
			beego.NSInclude(
				&controllers.TipoTitulacionController{},
			),
		),

		beego.NSNamespace("/titulacion",
			beego.NSInclude(
				&controllers.TitulacionController{},
			),
		),

		beego.NSNamespace("/proyecto_academico_institucion",
			beego.NSInclude(
				&controllers.ProyectoAcademicoInstitucionController{},
			),
		),
		beego.NSNamespace("/tr_proyecto_academico",
			beego.NSInclude(
				&controllers.TrProyectoAcademicoController{},
			),
		),

		beego.NSNamespace("/proyecto_academico_rol_persona_dependecia",
			beego.NSInclude(
				&controllers.ProyectoAcademicoRolPersonaDependeciaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
