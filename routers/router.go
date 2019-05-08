// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/planesticud/formacion_academica_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/dato_adicional_formacion_academica",
			beego.NSInclude(
				&controllers.DatoAdicionalFormacionAcademicaController{},
			),
		),

		beego.NSNamespace("/formacion_academica",
			beego.NSInclude(
				&controllers.FormacionAcademicaController{},
			),
		),

		beego.NSNamespace("/soporte_formacion_academica",
			beego.NSInclude(
				&controllers.SoporteFormacionAcademicaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
