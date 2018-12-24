// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"memplus_service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/knowledge_bases",
			beego.NSInclude(
				&controllers.KnowledgeBasesController{},
			),
		),

		beego.NSNamespace("/trade",
			beego.NSInclude(
				&controllers.TradeController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/knowledge_relations",
			beego.NSInclude(
				&controllers.KnowledgeBasesController{},
			),
		),

		beego.NSNamespace("/klg_record",
			beego.NSInclude(
				&controllers.CardController{},
			),
		),

		beego.NSNamespace("/loops",
			beego.NSInclude(
				&controllers.LoopController{},
			),
		),

		beego.NSNamespace("/cards",
			beego.NSInclude(
				&controllers.CardController{},
			),
		),

		beego.NSNamespace("/klg_dir",
			beego.NSInclude(
				&controllers.KlgDirController{},
			),
		),

		beego.NSNamespace("/history",
			beego.NSInclude(
				&controllers.HistoryController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
