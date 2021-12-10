package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:EnfasisController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:MetodologiaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:NivelFormacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoEnfasisController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoInstitucionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:ProyectoAcademicoRolTerceroDependenciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:RegistroCalificadoAcreditacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoRegistroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TipoTitulacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TitulacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"],
        beego.ControllerComments{
            Method: "PutProyectoEnfasis",
            Router: "/enfasis/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"],
        beego.ControllerComments{
            Method: "PutInformacionBasica",
            Router: "/informacion_basica/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/proyecto_academico_crud/controllers:TrProyectoAcademicoController"],
        beego.ControllerComments{
            Method: "PutProyectoRegistro",
            Router: "/registro/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
