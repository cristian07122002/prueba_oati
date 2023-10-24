package routers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"],
		beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"],
		beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"],
		beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

	beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"],
		beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

			beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"],
			beego.ControllerComments{
				Method: "Post",
				Router: "/",
				AllowHTTPMethods: []string{"post"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})
	
		beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"],
			beego.ControllerComments{
				Method: "GetAll",
				Router: "/",
				AllowHTTPMethods: []string{"get"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})
	
		beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"],
			beego.ControllerComments{
				Method: "GetOne",
				Router: "/:id",
				AllowHTTPMethods: []string{"get"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})
	
		beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"],
			beego.ControllerComments{
				Method: "Put",
				Router: "/:id",
				AllowHTTPMethods: []string{"put"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})
	
		beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:DetallesTutorialController"],
			beego.ControllerComments{
				Method: "Delete",
				Router: "/:id",
				AllowHTTPMethods: []string{"delete"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})

				beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:TutorialesController"],
		beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

			beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"],
			beego.ControllerComments{
				Method: "Post",
				Router: "/",
				AllowHTTPMethods: []string{"post"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})
	
		beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"],
			beego.ControllerComments{
				Method: "GetAll",
				Router: "/",
				AllowHTTPMethods: []string{"get"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})
	
		beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"],
			beego.ControllerComments{
				Method: "GetOne",
				Router: "/:id",
				AllowHTTPMethods: []string{"get"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})
	
		beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"],
			beego.ControllerComments{
				Method: "Put",
				Router: "/:id",
				AllowHTTPMethods: []string{"put"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})
	
		beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["prueba_oati/controllers:UsuariosController"],
			beego.ControllerComments{
				Method: "Delete",
				Router: "/:id",
				AllowHTTPMethods: []string{"delete"},
				MethodParams: param.Make(),
				Filters: nil,
				Params: nil})
}