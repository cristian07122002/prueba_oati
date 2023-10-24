// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"prueba_oati/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/usuarios",
			beego.NSInclude(
				&controllers.UsuariosController{},
			),
		),

		beego.NSNamespace("/tutoriales",
			beego.NSInclude(
				&controllers.TutorialesController{},
			),
		),

		beego.NSNamespace("/detalles_tutorial",
			beego.NSInclude(
				&controllers.DetallesTutorialController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
